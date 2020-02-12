package v3

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/api/interfaces"
	"rsi.com/go-training/api/v3/pet"
	"rsi.com/go-training/api/v3/user"
)

type api struct {
	db          *sql.DB
	engine      *gin.Engine
	authHandler auth.Handler
}

func (a api) Register(prefix string) {
	api := a.engine.Group(prefix)
	{
		group := api.Group("/v3")
		{
			group.GET("/user", a.authHandler.CurrentUser)
			configurePetRoutes(a.db, group, a.authHandler)
			configureUserRoutes(a.db, group)

		}
	}
}

func configureUserRoutes(db *sql.DB, group *gin.RouterGroup) {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	group.GET("/user/:email", userHandler.GetByEmail)

}

func configurePetRoutes(db *sql.DB, group *gin.RouterGroup, authHandler auth.Handler) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.Use(authHandler.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleBasicUser}))

	group.GET("/pets", petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", authHandler.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Create)
	group.PUT("/pet/:id", authHandler.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Update)
	group.DELETE("/pet/:id", authHandler.RequiresAuth(auth.AuthProfile{RoleRequired: auth.RoleAdministrator}), petHandler.Delete)
}

func NewApi(db *sql.DB, engine *gin.Engine, authHandler auth.Handler) interfaces.Api {
	return &api{db, engine, authHandler}
}
