package dbclient

import (
	"context"
	"fmt"
	"log"
	"university-db-admin/internal/config"

	"github.com/jackc/pgx/v5"
)

func NewClientPG(cfg config.DatabaseConfig) *pgx.Conn {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	log.Printf("connecting to %s\n", dsn)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	return conn
}
