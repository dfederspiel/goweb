package data

import (
	"database/sql"
	"log"
	"rsi.com/go-training/models"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func CreateAnimal(a models.Animal) models.Animal {
	statement, _ := db.Prepare("insert into pets (name, age, legs) values (?,?,?)")
	defer statement.Close()
	result, _ := statement.Exec(a.Name, a.Age, a.Legs)
	a.Id, _ = result.LastInsertId()
	return a
}

func GetAnimal(id int64) models.Animal {
	var name string
	var age int
	var legs bool
	row := db.QueryRow("select id, name, age, legs from pets where id = ?", id)
	err := row.Scan(&id, &name, &age, &legs)
	if err != nil {
		log.Println(err)
	}
	return models.Animal{
		Id:   id,
		Name: name,
		Age:  age,
		Legs: legs,
	}
}

func GetAnimals() []models.Animal {
	rows, err := db.Query("select id, name, age, legs from pets")
	defer rows.Close()

	var pets = make([]models.Animal, 0)
	for rows.Next() {
		var id int64
		var name string
		var age int
		var legs bool
		err = rows.Scan(&id, &name, &age, &legs)
		if err != nil {
			break
		}
		pets = append(pets, models.Animal{
			Id:   id,
			Name: name,
			Age:  age,
			Legs: legs,
		})
	}
	return pets
}
