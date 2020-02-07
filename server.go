package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
	"net/url"
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

	engine.GET("/callback", func(context *gin.Context) {
		fmt.Println(context.Query("code"))

		formData := url.Values{
			"code":          {context.Query("code")},
			"client_id":     {"90445840135-99mhv65o8m5kt3n6v46h6k1c2ie0eum1.apps.googleusercontent.com"},
			"client_secret": {"DrCd3z9oJdemHscZdstuCblb"},
			"redirect_uri":  {"http://localhost:8080/callback"},
			"grant_type":    {"authorization_code"},
		}

		var authResponse AuthResponse
		response, _ := http.PostForm("https://oauth2.googleapis.com/token", formData)
		getJson(response, &authResponse)
		fmt.Println(authResponse)
		http.SetCookie(context.Writer, &http.Cookie{
			Name:     "token",
			Value:    authResponse.IdToken,
			Expires:  time.Now().Add(120 * time.Minute),
			HttpOnly: true,
		})

	})

	engine.GET("/welcome", gin.WrapF(Welcome))
	engine.GET("/refresh", gin.WrapF(Refresh))

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IdToken     string `json:"id_token"`
	State       string `json:"session_state"`
}

func getJson(r *http.Response, target interface{}) error {
	defer r.Body.Close()
	//responseData, _ := ioutil.ReadAll(r.Body)
	//responseString := string(responseData)
	//fmt.Println(responseString)
	return json.NewDecoder(r.Body).Decode(target)
}

func RegisterMiddleware(g *gin.Engine) {
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
