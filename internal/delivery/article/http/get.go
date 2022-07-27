package http

import (
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	paginationCommon "github.com/aryahmph/kumparan-assesment/common/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPArticleDelivery) list(c *gin.Context) {
	// Bind query
	var paginationParams httpCommon.PaginationQueryParams
	if err := c.BindQuery(&paginationParams); err != nil {
		return
	}
	if paginationParams.Limit == 0 {
		paginationParams.Limit = httpCommon.PaginationDefaultLimit
	}

	var articleParams httpCommon.ArticleQueryParams
	if err := c.BindQuery(&articleParams); err != nil {
		return
	}

	// Get data
	articles, err := d.articleUsecase.GetAll(c.Request.Context(), paginationParams, articleParams)
	if err != nil {
		c.Error(err)
		return
	}

	// Convert data to response
	articlesLen := len(articles)
	articleResponses := make([]httpCommon.Article, articlesLen)
	for i := 0; i < articlesLen; i++ {
		articleResponses[i] = d.mapArticleModelToResponse(articles[i])
	}

	// Build response
	var nextCursor string
	if articlesLen > 0 {
		nextCursor = paginationCommon.EncodeCursor(articles[articlesLen-1].CreatedAt, articles[articlesLen-1].ID)
	}

	response := httpCommon.Response{
		Metadata: httpCommon.Metadata{
			Limit: paginationParams.Limit,
			Count: articlesLen,
			Cursor: httpCommon.Cursor{
				Self: paginationParams.Cursor,
				Next: nextCursor,
			}},
		Data: articleResponses,
	}

	c.JSON(http.StatusOK, response)
}
