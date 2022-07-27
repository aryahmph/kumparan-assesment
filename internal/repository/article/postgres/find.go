package postgres

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/kumparan-assesment/common/error"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

func (repository postgresArticleRepository) FindAll(ctx context.Context, statement string, args []interface{}) (articles []articleModel.Article, err error) {
	// "articles.id", "title", "body", "articles.created_at", "authors.id", "authors.name"
	rows, err := repository.db.QueryContext(ctx, statement, args...)
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var article articleModel.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Body, &article.CreatedAt, &article.Author.ID, &article.Author.Name)
		if errors.Is(err, sql.ErrNoRows) {
			return articles, errorCommon.NewNotFoundError("article not found")
		}
		articles = append(articles, article)
	}
	return articles, err
}
