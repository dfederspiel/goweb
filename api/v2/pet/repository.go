package pet

import "database/sql"

type Repository interface {
	GetAll() (pets []*Pet, err error)
	GetById(id string) (*Pet, error)
	Create(pet *Pet) error
	Update(pet *Pet) error
	DeleteById(id string) error
}

type repository struct {
	db *sql.DB
}

func (repo repository) GetAll() (pets []*Pet, err error) {
	rows, err := repo.db.Query("select name, age from pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Pet
		err = rows.Scan(&p.Name, &p.Age)
		if err != nil {
			return nil, err
		}
		pets = append(pets, &p)
	}
	return
}

func (repo repository) GetById(id string) (*Pet, error) {
	panic("implement me")
}

func (repo repository) Create(pet *Pet) error {
	panic("implement me")
}

func (repo repository) Update(pet *Pet) error {
	panic("implement me")
}

func (repo repository) DeleteById(id string) error {
	panic("implement me")
}

// NewRepository returns a Repository interface of unexported type repository, which implements all methods of the interface.
// It is important to understand the sleight of hand here
func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
