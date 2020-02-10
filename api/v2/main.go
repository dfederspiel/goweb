package v2

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v2/auth"
	"rsi.com/go-training/api/v2/pet"
)

var (
	authRepo    auth.Repository
	authService auth.Service
	authHandler auth.Handler
)

func Register(db *sql.DB, engine *gin.Engine) {

	authRepo = auth.NewRespository(db)
	authService = auth.NewService(authRepo)
	authHandler = auth.NewHandler(authService)

	authService.RegisterOauthCallbackRoute(engine)

	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			ConfigurePetRoutes(db, group)
		}
	}
}

func ConfigurePetRoutes(db *sql.DB, group *gin.RouterGroup) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.Use(authService.RequiresAuth(auth.AuthProfile{auth.RoleBasicUser}))

	group.GET("/user", authHandler.Get)

	group.GET("/pets", petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Create)
	group.PUT("/pet/:id", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Update)
	group.DELETE("/pet/:id", authService.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Delete)
}
