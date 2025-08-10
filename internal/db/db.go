package db

import (
	"fmt"

	"github.com/egon89/fc-event-driven-arch/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	return sqlx.Connect("pgx", dsn)
}
