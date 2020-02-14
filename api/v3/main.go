package v3

import (
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/api/v3/pet"
	"rsi.com/go-training/api/v3/user"
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
)

type Api interface {
	ConfigureUserRoutes(repo user.Repository)
	ConfigurePetRoutes(repo interfaces.PetRepository)
}

type api struct {
	engine      gin.IRouter
	authHandler auth.Handler
}

func (a api) ConfigurePetRoutes(repo interfaces.PetRepository) {
	petService := pet.NewService(repo)
	petHandler := pet.NewHandler(petService)

	a.engine.GET("/pets", petHandler.Get)
	a.engine.GET("/pet/:id", petHandler.GetById)
	a.engine.POST("/pet", a.authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Create)
	a.engine.PUT("/pet/:id", a.authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Update)
	a.engine.DELETE("/pet/:id", a.authHandler.RequiresAuth(models.RoleAdministrator), petHandler.Delete)
}

func (a api) ConfigureUserRoutes(repo user.Repository) {
	userService := user.NewService(repo)
	userHandler := user.NewHandler(userService)

	a.engine.GET("/user", a.authHandler.CurrentUser)
	a.engine.GET("/user/:email", a.authHandler.RequiresAuth(models.RoleAdministrator), userHandler.GetByEmail)
}

func NewApi(engine gin.IRouter, authHandler auth.Handler) Api {
	group := engine.Group("/api/v3")
	group.Use(authHandler.RequiresAuth(models.RoleBasicUser))
	return &api{group, authHandler}
}
