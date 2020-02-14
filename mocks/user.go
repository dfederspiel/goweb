package mocks

import (
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
)

type testableUserRepository struct{}

func (t testableUserRepository) GetByEmail(email string) (models.User, error) {
	return models.User{
		ID:    "1",
		Name:  "Buddy",
		Email: "david@federnet.com",
		Role:  0,
	}, nil
}

func NewTestableUserRepository() interfaces.UserRepository {
	return &testableUserRepository{}
}
