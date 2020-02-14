package user

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
	r.engine.GET("/user", r.auth.CurrentUser)
	r.engine.GET("/user/:email", r.auth.RequiresAuth(models.RoleAdministrator), r.handler.GetByEmail)
}

func NewRouter(r *gin.Engine, h Handler, a auth.Handler) Router {
	return &router{r, h, a}
}
