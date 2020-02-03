package data

import (
	"fmt"
	"log"
	"rsi.com/go-training/data/models"
)

func GetAnimals() ([]models.Animal, error) {
	rows, err := db.Query("select id, name, age, legs from pets")
	if err != nil {
		return []models.Animal{}, err
	}
	defer rows.Close()

	var pets = make([]models.Animal, 0)
	for rows.Next() {
		var a models.Animal
		err = rows.Scan(&a.Id, &a.Name, &a.Age, &a.Legs)
		if err != nil {
			return []models.Animal{}, err
		}
		pets = append(pets, a)
	}
	return pets, nil
}

func GetAnimal(id int64) (models.Animal, error) {
	var a models.Animal
	row := db.QueryRow("select id, name, age, legs from pets where id = ?", id)
	err := row.Scan(&a.Id, &a.Name, &a.Age, &a.Legs)
	if err != nil {
		log.Println(err)
		return models.Animal{}, err
	}
	return a, nil
}

func CreateAnimal(a models.Animal) (models.Animal, error) {
	statement, _ := db.Prepare("insert into pets (name, age, legs) values (?,?,?)")
	defer statement.Close()
	result, err := statement.Exec(a.Name, a.Age, a.Legs)
	if err != nil {
		return a, err
	}
	a.Id, _ = result.LastInsertId()
	fmt.Println(a.Id)
	return a, nil
}

func UpdateAnimal(a models.Animal) error {
	statement, _ := db.Prepare("update pets set name=?, age=?, legs=? where id=?")
	defer statement.Close()
	result, err := statement.Exec(a.Name, a.Age, a.Legs, a.Id)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return nil
}

func DeleteAnimal(id int64) error {
	statement, _ := db.Prepare("delete from pets where id=?")
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return nil
}
