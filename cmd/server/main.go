package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aryahmph/kumparan-assesment/common/env"
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	dbCommon "github.com/aryahmph/kumparan-assesment/common/postgres"
	articleDelivery "github.com/aryahmph/kumparan-assesment/internal/delivery/article/http"
	articleRepo "github.com/aryahmph/kumparan-assesment/internal/repository/article/postgres"
	authorRepo "github.com/aryahmph/kumparan-assesment/internal/repository/author/postgres"
	articleUc "github.com/aryahmph/kumparan-assesment/internal/usecase/article"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := env.LoadConfig()
	db := dbCommon.NewPostgres(cfg.PostgresMigrationPath, cfg.PostgresURL)
	router := httpCommon.NewHTTPServer()

	router.Router.Use(httpCommon.MiddlewareErrorHandler())
	root := router.Router.Group("/api")

	authorRepository := authorRepo.NewPostgresAuthorRepository(db)

	articleRepository := articleRepo.NewPostgresArticleRepository(db)
	articleUsecase := articleUc.NewArticleUsecase(articleRepository, authorRepository)
	articleDelivery.NewHTTPArticleDelivery(root.Group("/articles"), articleUsecase)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err)
	}
}
