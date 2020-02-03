package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"rsi.com/go-training/data"
	"rsi.com/go-training/services"
)

func main() {
	InitializeEnvironmentVariables()
	InitializeDB()
	StartServer()
}

func InitializeDB() {
	data.InitDB(os.Getenv("DATABASE"))
}

func StartServer() {
	g := gin.Default()

	RegisterMiddleware(g)
	RegisterApi(g)

	_ = g.Run()
}

func RegisterApi(g *gin.Engine) {
	api := g.Group(os.Getenv("API"))
	{
		services.RegisterRoutes(api)
	}
}

func RegisterMiddleware(g *gin.Engine) {
	g.Use(static.Serve("/", static.LocalFile("./www/dist", true)))
	g.Use(func(c *gin.Context) {
		fmt.Println(c.Request)
	})
}

func InitializeEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
