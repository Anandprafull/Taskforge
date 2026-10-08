package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func CreateTables(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		command TEXT NOT NULL
	);
	`
	_, err := conn.Exec(context.Background(), query)

	return err
}