package services

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

type Animal struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Legs bool   `json:"legs"`
}

/*StartServer sets up and runs our web server
 */
func RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/animals", GetAnimals())
	router.GET("/animal/:id", GetAnimal())
	router.POST("/animal", CreateAnimal())
	router.PUT("/animal/:id", UpdateAnimal())
	router.DELETE("/animal/:id", DeleteAnimal())
}

func CreateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		var a Animal
		_ = c.BindJSON(&a)

		db, err := sql.Open("sqlite3", "./db/goweb.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		statement, _ := db.Prepare("insert into pets (name, age, legs) values (?,?,?)")
		defer statement.Close()
		result, _ := statement.Exec(a.Name, a.Age, a.Legs)
		a.Id, _ = result.LastInsertId()

		c.JSON(http.StatusOK, a)
	}
}

func GetAnimals() func(c *gin.Context) {
	return func(c *gin.Context) {

		db, err := sql.Open("sqlite3", "./db/goweb.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, _ := db.Query("select * from pets")
		defer rows.Close()

		var pets = make([]Animal, 0)
		for rows.Next() {
			var id int64
			var name string
			var age int
			var legs bool
			err = rows.Scan(&id, &name, &age, &legs)
			if err != nil {
				break
			}
			pets = append(pets, Animal{
				Id:   id,
				Name: name,
				Age:  age,
				Legs: legs,
			})
		}

		c.JSON(http.StatusOK, pets)
	}
}

func GetAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {

		db, err := sql.Open("sqlite3", "./db/goweb.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var id int64
		var name string
		var age int
		var legs bool
		row := db.QueryRow("select * from pets where id = :id", c.Param("id"))
		row.Scan(&id, &name, &age, &legs)

		c.JSON(http.StatusOK, Animal{
			Id:   id,
			Name: name,
			Age:  age,
			Legs: legs,
		})
	}
}

func UpdateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {

		var a Animal
		_ = c.BindJSON(&a)

		db, err := sql.Open("sqlite3", "./db/goweb.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		statement, _ := db.Prepare("update pets set name=?, age=?, legs=? where id=?")
		defer statement.Close()

		statement.Exec(a.Name, a.Age, a.Legs, c.Param("id"))

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		c.JSON(http.StatusOK, Animal{
			Id:   id,
			Name: a.Name,
			Age:  a.Age,
			Legs: a.Legs,
		})
	}
}

func DeleteAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "./db/goweb.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		statement, _ := db.Prepare("delete from pets where id=?")
		defer statement.Close()

		statement.Exec(c.Param("id"))

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		c.JSON(http.StatusOK, id)
	}
}
