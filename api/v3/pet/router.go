package pet

import (
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/models"
)

type Router interface {
	Configure()
	Engine() *gin.Engine
}

type router struct {
	engine  *gin.Engine
	handler Handler
	auth    auth.Handler
}

func (r router) Engine() *gin.Engine {
	return r.engine
}

func (r router) Configure() {
	r.engine.GET("/pets", r.handler.Get)
	r.engine.GET("/pet/:id", r.handler.GetById)
	r.engine.POST("/pet", r.auth.RequiresAuth(models.RoleAdministrator), r.handler.Create)
	r.engine.PUT("/pet/:id", r.auth.RequiresAuth(models.RoleAdministrator), r.handler.Update)
	r.engine.DELETE("/pet/:id", r.auth.RequiresAuth(models.RoleAdministrator), r.handler.Delete)
}

func NewRouter(r *gin.Engine, h Handler, a auth.Handler) Router {
	return &router{r, h, a}
}
