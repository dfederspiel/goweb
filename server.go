package main

import (
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/api/v1"
	"rsi.com/go-training/api/v2"
	"rsi.com/go-training/api/v3"
	"rsi.com/go-training/api/v3/pet"
	"rsi.com/go-training/api/v3/user"
	"rsi.com/go-training/data"
	"time"
)

var engine *gin.Engine
var db *sql.DB

func initializeDB(dataSource string) {
	database, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Panic(err)
	}
	db = database

	seeder := data.NewSeeder(db)
	seeder.Seed()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	v1.ConfigureDB(db)
}

func startServer() {
	engine = gin.Default()
	registerMiddleware(engine)

	authHandler := configureOauth()

	v1.Register(engine)
	v2.Register(db, engine, authHandler)

	configureV3Api(authHandler)

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

func configureV3Api(authHandler auth.Handler) {
	api := v3.NewApi(engine, authHandler)
	api.ConfigurePetRoutes(pet.NewRepository(db))
	api.ConfigureUserRoutes(user.NewRepository(db))
}

func configureOauth() auth.Handler {
	authRepo := auth.NewRespository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	engine.GET("/callback", authHandler.Callback)
	engine.POST("/logout", authHandler.Logout)

	return authHandler
}

func registerMiddleware(g *gin.Engine) {
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
