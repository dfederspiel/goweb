package pet

import (
	"database/sql"
	"fmt"
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
	"strconv"
)

type repository struct {
	db *sql.DB
}

func (repo repository) GetAll() (pets []*models.Pet, err error) {
	rows, _ := repo.db.Query("select id, name, age, color from pets")
	defer rows.Close()

	for rows.Next() {
		var p models.Pet
		rows.Scan(&p.ID, &p.Name, &p.Age, &p.Color)
		pets = append(pets, &p)
	}
	return
}

func (repo repository) GetById(id string) (pet *models.Pet, err error) {
	row := repo.db.QueryRow("select id, name, age, color from pets where id = ?", id)
	pet = &models.Pet{}
	row.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Color)
	return
}

func (repo repository) Create(pet *models.Pet) (err error) {
	statement, _ := repo.db.Prepare("insert into pets (name, age, color, legs) values (?,?,?,?)")
	defer statement.Close()
	result, _ := statement.Exec(pet.Name, pet.Age, pet.Color, pet.Legs)
	id, err := result.LastInsertId()
	pet.ID = strconv.Itoa(int(id))
	fmt.Println(pet.ID)
	return nil
}

func (repo repository) Update(pet *models.Pet) (err error) {
	statement, _ := repo.db.Prepare("update pets set name=?, age=?, legs=?, color=? where id=?")
	defer statement.Close()
	result, _ := statement.Exec(pet.Name, pet.Age, pet.Legs, pet.Color, pet.ID)
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return
}

func (repo repository) DeleteById(id string) (err error) {
	statement, _ := repo.db.Prepare("delete from pets where id=?")
	defer statement.Close()
	result, _ := statement.Exec(id)
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return
}

// NewRepository returns a Repository interface of unexported type repository, which implements all methods of the interface.
// It is important to understand the sleight of hand here
func NewRepository(db *sql.DB) interfaces.PetRepository {
	return &repository{db}
}
