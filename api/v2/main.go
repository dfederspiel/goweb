package v2

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v2/pet"
)

func Register(db *sql.DB, engine *gin.Engine) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			group.GET("/pets", petHandler.Get)
			group.GET("/pet/:id", petHandler.GetById)
			group.POST("/pet", petHandler.Create)
			group.PUT("/pet/:id", petHandler.Update)
			group.DELETE("/pet/:id", petHandler.Delete)
		}
	}
}
