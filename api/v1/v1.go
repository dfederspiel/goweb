package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v1/pet"
)

var db *sql.DB

func ConfigureDB(database *sql.DB) {
	db = database
	pet.ConfigureDB(db)
}

func Register(g *gin.Engine) {
	api := g.Group(os.Getenv("API"))
	{
		group := api.Group("/v1")
		{
			pet.Routes(group)
		}
	}
}
