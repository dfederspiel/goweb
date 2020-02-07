package v2

import (
	"database/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v2/pet"
)

var middleware *jwt.GinJWTMiddleware

func Register(db *sql.DB, engine *gin.Engine, auth *jwt.GinJWTMiddleware) {
	middleware = auth
	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			ConfigurePetRoutes(db, group)
			//ConfigureOwnerRoutes
		}
	}
}

func ConfigurePetRoutes(db *sql.DB, group *gin.RouterGroup) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.Use(middleware.MiddlewareFunc())

	group.GET("/pets", petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", petHandler.Create)
	group.PUT("/pet/:id", petHandler.Update)
	group.DELETE("/pet/:id", petHandler.Delete)
}
