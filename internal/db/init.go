package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {

	urlDB := BuildPostgreDSN()

	db, err := sql.Open("postgres", urlDB)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("data base ping error: %w", errPing)
	}
	return db, nil
}
