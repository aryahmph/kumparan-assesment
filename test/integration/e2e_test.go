package integration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aryahmph/kumparan-assesment/common/env"
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	dbCommon "github.com/aryahmph/kumparan-assesment/common/postgres"
	articleDelivery "github.com/aryahmph/kumparan-assesment/internal/delivery/article/http"
	articleRepo "github.com/aryahmph/kumparan-assesment/internal/repository/article/postgres"
	authorRepo "github.com/aryahmph/kumparan-assesment/internal/repository/author/postgres"
	articleUc "github.com/aryahmph/kumparan-assesment/internal/usecase/article"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	port        int
	db          *sql.DB
	dbMigration *migrate.Migrate
	handler     http.Handler
}

func (suite *ExampleTestSuite) SetupSuite() {
	cfg := env.LoadConfig()
	m, err := migrate.New(cfg.PostgresMigrationPath, cfg.PostgresURL)
	suite.Require().NoError(err)

	suite.port = cfg.Port
	suite.dbMigration = m
	suite.db = dbCommon.NewPostgres(cfg.PostgresMigrationPath, cfg.PostgresURL)

	router := httpCommon.NewHTTPServer()
	router.Router.Use(httpCommon.MiddlewareErrorHandler())
	root := router.Router.Group("/api")
	authorRepository := authorRepo.NewPostgresAuthorRepository(suite.db)
	articleRepository := articleRepo.NewPostgresArticleRepository(suite.db)
	articleUsecase := articleUc.NewArticleUsecase(articleRepository, authorRepository)
	articleDelivery.NewHTTPArticleDelivery(root.Group("/articles"), articleUsecase)
	suite.handler = router.Router
}

func (suite *ExampleTestSuite) TearDownSuite() {
	suite.NoError(suite.dbMigration.Down())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (suite *ExampleTestSuite) Test_EndToEnd_CreateArticle_Success() {
	requestBody := `
		{
			"title": "Shinta Kamdani Bawa Perempuan Indonesia Berdaya Ekonomi di Forum B20",
			"body": "Forum Business of 20 atau B20 yang diketuai oleh Shinta Widjaja Kamdani, akan menggagas sejumlah agenda dalam Presidensi Indonesia di G20. Shinta Kamdani menjadi perempuan pertama asal Asia yang memimpin salah satu forum bisnis terbesar di dunia ini.",
			"author_id": "d3ad2722-675e-4d94-935a-a31fb5495dbf"
		}
	`
	request := httptest.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/articles", suite.port),
		strings.NewReader(requestBody),
	)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	suite.handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	id := responseBody["data"].(map[string]interface{})["id"].(string)
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.NotEmpty(id)

	parse, err := uuid.Parse(id)
	suite.Require().NoError(err)
	suite.Equal("VERSION_4", parse.Version().String())
}

func (suite *ExampleTestSuite) Test_EndToEnd_CreateArticle_Failed() {
	requestBody := `
		{
			"title": "Shinta Kamdani Bawa Perempuan Indonesia Berdaya Ekonomi di Forum B20",
			"author_id": "d3ad2722-675e-4d94-935a-a31fb5495dbf"
		}
	`
	request := httptest.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/api/articles", suite.port),
		strings.NewReader(requestBody),
	)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	suite.handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal(http.StatusBadRequest, int(responseBody["code"].(float64)))
	suite.NotEmpty(responseBody["message"].(string))
	suite.Equal("required", responseBody["errors"].(map[string]interface{})["body"].(string))
}

func (suite *ExampleTestSuite) Test_EndToEnd_ListArticle() {
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/api/articles", suite.port), nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	suite.handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var responseBody httpCommon.Response
	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(10, responseBody.Metadata.Limit)
	suite.Empty(responseBody.Metadata.Cursor.Self)
	suite.Equal(len(responseBody.Data.([]interface{})), responseBody.Metadata.Count)
}

func (suite *ExampleTestSuite) Test_EndToEnd_ListArticle_Query() {
	queryLimit := 5
	queryKeyword := "apple"
	queryAuthor := "arya"
	params := url.Values{}
	params.Add("limit", strconv.Itoa(queryLimit))
	params.Add("author", queryAuthor)
	params.Add("query", queryKeyword)

	request := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:%d/api/articles?%s", suite.port, params.Encode()),
		nil,
	)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	suite.handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	suite.Require().NoError(err)

	var responseBody httpCommon.Response
	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	err = json.Unmarshal(body, &responseBody)
	suite.Require().NoError(err)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(queryLimit, responseBody.Metadata.Limit)
	suite.Empty(responseBody.Metadata.Cursor.Self)
	suite.Equal(len(responseBody.Data.([]interface{})), responseBody.Metadata.Count)
}
