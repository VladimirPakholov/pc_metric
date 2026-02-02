package db

import (
	"database/sql"
	"fmt"
	"pc_metric/internal/db/repository"

	_ "github.com/lib/pq"
)

func InitDB() (*repository.Repository, error) {

	urlDB := BuildPostgreDSN()

	db, err := sql.Open("postgres", urlDB)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("data base ping error: %w", errPing)
	}
	fmt.Println("Connected to data base - successfull!")

	return &repository.Repository{DB: db}, nil

}
