package dbclient

import (
	"context"
	"fmt"
	"log"
	"university-db-admin/internal/config"

	"github.com/jackc/pgx/v5"
)

func NewClient(db_cfg config.DatabaseConfig) *pgx.Conn {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db_cfg.User, db_cfg.Password, db_cfg.Host, db_cfg.Port, db_cfg.Name)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)

	return conn
}
