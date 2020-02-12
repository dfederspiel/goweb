package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	GetByEmail(c *gin.Context)
}

type handler struct {
	service Service
}

func (h handler) GetByEmail(c *gin.Context) {
	user, _ := h.service.GetByEmail(c.Param("email"))
	c.JSON(http.StatusOK, user)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}
