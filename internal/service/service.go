package service

import (
	"errors"
	"time"
)

type UserService struct {
	repo MetricRepository
}

func NewUserService(repo MetricRepository) *UserService {
	return &UserService{repo: repo}
}
func (u *UserService) AddMetric(createdAt time.Time, message string) error {
	if message == "" {
		return errors.New("message is empty")
	}
	return u.repo.AddMetric(time.Now(), message)
}
