package services

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"rsi.com/go-training/data"
	"rsi.com/go-training/models"
	"strconv"
)

const dbLocation = "./data/goweb.db"

func RegisterAnimalRoutes(router *gin.RouterGroup) {

	router.GET("/animals", GetAnimals())
	router.GET("/animal/:id", GetAnimal())
	router.POST("/animal", CreateAnimal())
	router.PUT("/animal/:id", UpdateAnimal())
	router.DELETE("/animal/:id", DeleteAnimal())
}

func CreateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		var a models.Animal
		_ = c.BindJSON(&a)

		a = data.CreateAnimal(a)

		c.JSON(http.StatusOK, a)
	}
}

func GetAnimals() func(c *gin.Context) {
	return func(c *gin.Context) {

		animals := data.GetAnimals()
		c.JSON(http.StatusOK, animals)
	}
}

func GetAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		animal := data.GetAnimal(id)
		c.JSON(http.StatusOK, animal)
	}
}

func UpdateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {

		var a models.Animal
		_ = c.BindJSON(&a)

		db, err := sql.Open("sqlite3", dbLocation)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		statement, _ := db.Prepare("update pets set name=?, age=?, legs=? where id=?")
		defer statement.Close()

		statement.Exec(a.Name, a.Age, a.Legs, c.Param("id"))

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		c.JSON(http.StatusOK, models.Animal{
			Id:   id,
			Name: a.Name,
			Age:  a.Age,
			Legs: a.Legs,
		})
	}
}

func DeleteAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", dbLocation)
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
