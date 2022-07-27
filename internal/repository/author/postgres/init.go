package postgres

import (
	"database/sql"
	authorRepo "github.com/aryahmph/kumparan-assesment/internal/repository/author"
)

type postgresAuthorRepository struct {
	db *sql.DB
}

func NewPostgresAuthorRepository(db *sql.DB) authorRepo.Repository {
	return postgresAuthorRepository{
		db: db,
	}
}
