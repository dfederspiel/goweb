package services

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	RegisterAnimalRoutes(router)
}
