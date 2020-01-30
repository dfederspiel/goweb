package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Animal struct {
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

		a.Name += " (Updated)"
		c.JSON(200, a)
	}
}

func GetAnimals() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Get  all animals")
	}
}

func GetAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, _ = fmt.Fprintf(c.Writer, "Get animal %v", c.Param("id"))
	}
}

func UpdateAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, _ = fmt.Fprintf(c.Writer, "Update animal %v", c.Param("id"))
	}
}

func DeleteAnimal() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, _ = fmt.Fprintf(c.Writer, "Delete animal %v", c.Param("id"))
	}
}
