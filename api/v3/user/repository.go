package user

import (
	"database/sql"
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
)

type Repository interface {
	GetByEmail(email string) (models.User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) GetByEmail(email string) (models.User, error) {
	row := r.db.QueryRow("select id, name, role, email from users where email = ?", email)
	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Role, &u.Email)
	if err != nil {
		return u, err
	}
	return u, nil
}

func NewRepository(db *sql.DB) interfaces.UserRepository {
	return &repository{db}
}
