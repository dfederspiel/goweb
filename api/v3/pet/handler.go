package pet

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	service Service
}

func (h handler) Get(c *gin.Context) {
	pets, _ := h.service.GetAll()
	c.JSON(http.StatusOK, pets)
}

func (h handler) GetById(c *gin.Context) {
	pet, _ := h.service.GetById(c.Param("id"))
	c.JSON(http.StatusOK, pet)
}

func (h handler) Create(c *gin.Context) {
	p := Pet{}
	err := c.Bind(&p)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	err = h.service.Create(&p)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, p)
}

func (h handler) Update(c *gin.Context) {
	var p Pet
	p.ID = c.Param("id")
	_ = c.BindJSON(&p)

	err := h.service.Update(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.Status(http.StatusOK)
	}
}

func (h handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.Status(http.StatusOK)
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}
