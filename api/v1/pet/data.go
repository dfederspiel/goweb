package pet

import (
	"fmt"
	"log"
)

func GetPets() (pets []Pet, err error) {

	rows, err := db.Query("select id, name, age, legs from pets")
	if err != nil {
		return []Pet{}, err
	}
	defer rows.Close()

	pets = make([]Pet, 0)
	for rows.Next() {
		var a Pet
		err = rows.Scan(&a.Id, &a.Name, &a.Age, &a.Legs)
		if err != nil {
			return []Pet{}, err
		}
		pets = append(pets, a)
	}
	return
}

func GetPet(id int64) (a Pet, err error) {
	a = Pet{}
	row := db.QueryRow("select id, name, age, legs from pets where id = ?", id)
	err = row.Scan(&a.Id, &a.Name, &a.Age, &a.Legs)
	if err != nil {
		log.Println(err)
	}
	return
}

func CreatePet(a Pet) (Pet, error) {
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

func UpdatePet(a Pet) (err error) {
	statement, _ := db.Prepare("update pets set name=?, age=?, legs=? where id=?")
	defer statement.Close()
	result, err := statement.Exec(a.Name, a.Age, a.Legs, a.Id)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return
}

func DeletePet(id int64) (err error) {
	statement, _ := db.Prepare("delete from pets where id=?")
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected()
	fmt.Println(i)
	return
}
