package postgres

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/kumparan-assesment/common/error"
	authorModel "github.com/aryahmph/kumparan-assesment/internal/model/author"
)

func (repository postgresAuthorRepository) FindByID(ctx context.Context, id string) (author authorModel.Author, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, name, created_at FROM authors WHERE id = $1", id)
	err = row.Scan(&author.ID, &author.Name, &author.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return author, errorCommon.NewNotFoundError("author not found")
	}
	return author, err
}
