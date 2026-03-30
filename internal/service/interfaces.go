package service

import "time"

type MetricRepository interface {
	AddMetric(timestamp time.Time, message string) error
}
