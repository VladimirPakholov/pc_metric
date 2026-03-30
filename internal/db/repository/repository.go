package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// add units (create table)
func (r *Repository) AddMetric(createdAt time.Time, message string) error {
	_, err := r.db.Exec(
		`INSERT INTO logs_metric (created_at, message)
 		VALUES ($1, $2)`,
		createdAt,
		message,
	)
	if err != nil {
		return fmt.Errorf("insert log error: %w", err)
	}
	return nil
}
