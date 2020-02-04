package v2

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/pets", HandleGetPets())
	router.GET("/pet/:id", HandleGetPet())
	router.POST("/pet", HandleCreatePet())
	router.PUT("/pet/:id", HandleUpdatePet())
	router.DELETE("/pet/:id", HandleDeletePet())
}
