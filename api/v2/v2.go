package v2

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func ConfigureDB(database *sql.DB) {
	db = database
}

func Register(router *gin.RouterGroup) {
	group := router.Group("/v2")
	{
		Routes(group)
	}
}
