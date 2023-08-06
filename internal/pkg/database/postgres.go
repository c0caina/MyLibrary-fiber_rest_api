package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func newPostgreSQL() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_SERVER_URL"))
	if err != nil {
		defer conn.Close(context.Background())
		return conn, err
	}
	return conn, err
}
