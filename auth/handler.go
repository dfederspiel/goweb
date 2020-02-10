package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	Get(c *gin.Context)
}

type handler struct {
	service Service
}

func (h handler) Get(c *gin.Context) {
	user, _ := h.service.Get(c)
	c.JSON(http.StatusOK, user)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}
