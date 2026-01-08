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


	//migration code
	migrate(conn, "migration/user.sql")
	migrate(conn, "migration/post.sql")
	return conn
}

func migrate(pool *pgxpool.Pool, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, string(sqlBytes))
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
