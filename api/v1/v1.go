package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/api/v1/pet"
)

var db *sql.DB

func ConfigureDB(database *sql.DB) {
	db = database
	pet.ConfigureDB(db)
}

func Register(router *gin.RouterGroup) {
	group := router.Group("/v1")
	{
		pet.Routes(group)
	}
}
