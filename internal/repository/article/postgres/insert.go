package postgres

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/kumparan-assesment/common/error"
	aModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

func (repository postgresArticleRepository) Insert(ctx context.Context, article aModel.Article) (id string, err error) {
	row := repository.db.QueryRowContext(ctx,
		"INSERT INTO articles(author_id, title, body) VALUES ($1,$2,$3) RETURNING id;",
		article.Author.ID, article.Title, article.Body,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("article not found")
	}
	return id, err
}
