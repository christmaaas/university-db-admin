package dbclient

import (
	"context"
	"fmt"
	"log"
	"university-db-admin/internal/config"

	"github.com/jackc/pgx/v5"
)

func NewClientPG(db_cfg config.DatabaseConfig) *pgx.Conn {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db_cfg.User, db_cfg.Password, db_cfg.Host, db_cfg.Port, db_cfg.Name)

	log.Printf("connecting to %s\n", dsn)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	return conn
}
