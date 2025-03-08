package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgconn"
)

func handlePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
		log.Println(newErr)
		return newErr
	}
	return err
}
