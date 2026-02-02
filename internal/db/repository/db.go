package repository

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func (d *Repository) AddMetricDB(createdAt time.Time, message string) error {
	_, err := d.DB.Exec(
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
