package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"os"
	"rsi.com/go-training/data"
	"rsi.com/go-training/services"
	"time"
)

func main() {
	initializeEnvironmentVariables()
	initializeDB()
	startServer()
}

func initializeDB() {
	db := initDB(os.Getenv("DATABASE"))
	data.ConfigureDB(db)
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

	registerMiddleware(g)
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

func registerMiddleware(g *gin.Engine) {
	g.Use(static.Serve("/", static.LocalFile("./www", true)))
	g.Use(func(c *gin.Context) {
		fmt.Println(c.Request)
	})
	g.Use(stats.RequestStats())
	g.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

func initializeEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
