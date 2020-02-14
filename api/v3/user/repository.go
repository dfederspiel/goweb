package user

import "database/sql"

type Repository interface {
	GetByEmail(email string) (User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) GetByEmail(email string) (User, error) {
	row := r.db.QueryRow("select id, name, role, email from users where email = ?", email)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Role, &u.Email)
	if err != nil {
		return u, err
	}
	return u, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
