package pet

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rsi.com/go-training/models"
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
	p := models.Pet{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		h.service.Create(&p)
		c.JSON(http.StatusOK, p)

	}
}

func (h handler) Update(c *gin.Context) {
	var p models.Pet
	p.ID = c.Param("id")
	_ = c.BindJSON(&p)

	h.service.Update(&p)
	c.Status(http.StatusOK)
}

func (h handler) Delete(c *gin.Context) {
	id := c.Param("id")
	h.service.DeleteById(id)
	c.Status(http.StatusOK)
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}
