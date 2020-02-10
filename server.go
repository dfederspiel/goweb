package main

import (
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"rsi.com/go-training/api/v1"
	"rsi.com/go-training/api/v2"
	"time"
)

var engine *gin.Engine
var db *sql.DB

func initializeDB(dataSource string) {
	d, err := sql.Open("sqlite3", dataSource)
	db = d
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	v1.ConfigureDB(db)
}

func startServer() {
	engine = gin.Default()

	RegisterMiddleware(engine)

	v1.Register(engine)
	v2.Register(db, engine)

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

func RegisterMiddleware(g *gin.Engine) {
	configureStaticDirectoryMiddleware(g)
	configureStatsMiddleware(g)
	configureCORSMiddleware(g)
}

func configureCORSMiddleware(g *gin.Engine) {
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

func configureStaticDirectoryMiddleware(g *gin.Engine) {
	g.Use(static.Serve("/", static.LocalFile("./www", true)))
}

func configureStatsMiddleware(g *gin.Engine) {
	g.Use(stats.RequestStats())
	g.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}
