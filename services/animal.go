package services

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"rsi.com/go-training/data"
	"rsi.com/go-training/models"
	"strconv"
)

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

		a, err := data.CreateAnimal(a)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusCreated, a)
		}
	}
}

func GetAnimals() func(c *gin.Context) {
	return func(c *gin.Context) {

		animals, err := data.GetAnimals()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, animals)
	}
}

func GetAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		animal, err := data.GetAnimal(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, animal)
		}
	}
}

func UpdateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {

		var a models.Animal
		_ = c.BindJSON(&a)

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		a.Id = id

		err := data.UpdateAnimal(a)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func DeleteAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := data.DeleteAnimal(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.Status(http.StatusOK)
	}
}
