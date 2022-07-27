package postgres

import (
	"database/sql"
	articleRepo "github.com/aryahmph/kumparan-assesment/internal/repository/article"
)

type postgresArticleRepository struct {
	db *sql.DB
}

func NewPostgresArticleRepository(db *sql.DB) articleRepo.Repository {
	return postgresArticleRepository{
		db: db,
	}
}
