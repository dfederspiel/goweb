package user

import "rsi.com/go-training/models"

type Service interface {
	GetByEmail(email string) (models.User, error)
}

type service struct {
	repo Repository
}

func (s service) GetByEmail(email string) (models.User, error) {
	return s.repo.GetByEmail(email)
}

func NewService(repo Repository) Service {
	return &service{repo}
}
