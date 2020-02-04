package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleCreatePet() func(c *gin.Context) {
	return func(c *gin.Context) {
		var a Pet
		_ = c.BindJSON(&a)
		a, err := CreatePet(a)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusCreated, a)
		}
	}
}

func HandleGetPets() func(c *gin.Context) {
	return func(c *gin.Context) {

		Pets, err := GetPets()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, Pets)
	}
}

func HandleGetPet() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		Pet, err := GetPet(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, Pet)
		}
	}
}

func HandleUpdatePet() func(c *gin.Context) {
	return func(c *gin.Context) {

		var a Pet
		_ = c.BindJSON(&a)

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		a.Id = id

		err := UpdatePet(a)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func HandleDeletePet() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := DeletePet(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.Status(http.StatusOK)
	}
}
