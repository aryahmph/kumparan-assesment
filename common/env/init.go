package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port                  int    `env:"PORT,unset" envDefault:"8000"`
	PostgresURL           string `env:"POSTGRESQL_CONNECTION_URL,unset"`
	PostgresMigrationPath string `env:"POSTGRES_MIGRATION_PATH,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
