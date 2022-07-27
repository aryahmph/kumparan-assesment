package postgres

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func NewPostgres(migrationPath, connectionURL string) *sql.DB {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatalln("failed to create connection to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("failed to connect database:", err)
	}

	migrateDB(migrationPath, connectionURL)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func migrateDB(migrationPath, dsn string) {
	migration, err := migrate.New(migrationPath, dsn)
	if err != nil {
		log.Fatalln("failed to create migration", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln("failed to migrate", err)
	}
}
