package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"os"
	"rsi.com/go-training/data"
	"rsi.com/go-training/services"
)

func main() {
	initializeEnvironmentVariables()
	initializeDB()
	startServer()
}

func initializeDB() {
	db := initDB(os.Getenv("DATABASE"))
	data.InitDB(db)
}

func initDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}

func startServer() {
	g := gin.Default()

	registerMiddlewares(g)
	registerApi(g)
	g.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	_ = g.Run()
}

func registerApi(g *gin.Engine) {
	api := g.Group(os.Getenv("API"))
	{
		services.RegisterRoutes(api)
	}
}

func registerMiddlewares(g *gin.Engine) {
	g.Use(static.Serve("/", static.LocalFile("./www/dist", true)))
	g.Use(func(c *gin.Context) {
		fmt.Println(c.Request)
	})
	g.Use(stats.RequestStats())
}

func initializeEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
