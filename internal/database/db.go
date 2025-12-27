package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLmode  string
}

func NewDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBname,
		cfg.SSLmode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("data base ping error: %w", err)
	}

	fmt.Println("Connected to data base - successfull!")

	return db, nil
}

func NewTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS logs_metric (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL,
	message TEXT NOT NULL
);`)
	return err
}

func InsertLogMetric(db *sql.DB, createdAt time.Time, message string) error {
	_, err := db.Exec(
		`INSERT INTO logs_metric (created_at, message)
		VALUES ($1, $2)`,
		createdAt,
		message,
	)
	if err != nil {
		return fmt.Errorf("Insert log error: %w", err)
	}
	return nil
}
