package v2

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v2/auth"
	"rsi.com/go-training/api/v2/pet"
	"rsi.com/go-training/api/v2/user"
)

func Register(db *sql.DB, engine *gin.Engine) {

	authRepo := auth.NewRespository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	engine.GET("/callback", authHandler.Callback)

	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			group.GET("/user", authHandler.CurrentUser)
			ConfigurePetRoutes(db, group, authService)
			ConfigureUserRoutes(db, group)
		}
	}
}

func ConfigureUserRoutes(db *sql.DB, group *gin.RouterGroup) {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	group.GET("/user/:email", userHandler.GetByEmail)

}

func ConfigurePetRoutes(db *sql.DB, group *gin.RouterGroup, authService auth.Service) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.Use(authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleBasicUser}))

	group.GET("/pets", petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Create)
	group.PUT("/pet/:id", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Update)
	group.DELETE("/pet/:id", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Delete)
}
