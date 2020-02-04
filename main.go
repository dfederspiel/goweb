package main

import (
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"os"
	"rsi.com/go-training/api/v1"
	"rsi.com/go-training/api/v2"
	"time"
)

var engine *gin.Engine

func main() {
	initializeEnvironmentVariables()
	initializeDB(os.Getenv("DATABASE"))
	startServer()
}

func initializeEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initializeDB(dataSource string) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	setAPIDataContext(db)
}

func setAPIDataContext(db *sql.DB) {
	v1.ConfigureDB(db)
	v2.ConfigureDB(db)
}

func startServer() {
	engine = gin.Default()

	registerMiddleware(engine)
	registerApi(engine)

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

func registerApi(g *gin.Engine) {
	api := g.Group(os.Getenv("API"))
	{
		v1.Register(api)
		v2.Register(api)
	}
}

func registerMiddleware(g *gin.Engine) {
	configureStaticDirectoryMiddleware(g)
	configureStatsMiddleware(g)
	configureCORSMiddleware(g)
}

func configureCORSMiddleware(g *gin.Engine) gin.IRoutes {
	return g.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

func configureStaticDirectoryMiddleware(g *gin.Engine) gin.IRoutes {
	return g.Use(static.Serve("/", static.LocalFile("./www", true)))
}

func configureStatsMiddleware(g *gin.Engine) {
	g.Use(stats.RequestStats())
	g.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}
