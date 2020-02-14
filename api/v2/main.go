package v2

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/api/v2/pet"
	"rsi.com/go-training/api/v2/user"
	"rsi.com/go-training/models"
)

func Register(db *sql.DB, engine *gin.Engine, authHandler auth.Handler) {
	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			group.GET("/user", authHandler.CurrentUser)
			ConfigurePetRoutes(db, group, authHandler)
			ConfigureUserRoutes(db, group, authHandler)
		}
	}
}

func ConfigureUserRoutes(db *sql.DB, group *gin.RouterGroup, authHandler auth.Handler) {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	group.GET("/user/:email", userHandler.GetByEmail)

}

func ConfigurePetRoutes(db *sql.DB, group *gin.RouterGroup, authHandler auth.Handler) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.Use(authHandler.RequiresAuth(models.RoleBasicUser))

	group.GET("/pets", petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Create)
	group.PUT("/pet/:id", authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Update)
	group.DELETE("/pet/:id", authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Delete)
}
