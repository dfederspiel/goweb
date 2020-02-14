package interfaces

import "rsi.com/go-training/models"

type UserRepository interface {
	GetByEmail(email string) (models.User, error)
}
