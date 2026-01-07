package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		print("Unable to connect db: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	return conn
}
