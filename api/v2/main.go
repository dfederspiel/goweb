package v2

import (
	"database/sql"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"rsi.com/go-training/api/v2/pet"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("AIzaSyDttqL8yqdk2tBjW6tJki5s_uVPf3jfYP8"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		c.Request.Header.Set("Authorization", "Bearer "+cookie)
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
	}
}
func Register(db *sql.DB, engine *gin.Engine) {
	api := engine.Group(os.Getenv("API"))
	{
		group := api.Group("/v2")
		{
			ConfigurePetRoutes(db, group)
			//ConfigureOwnerRoutes
		}
	}
}

func ConfigurePetRoutes(db *sql.DB, group *gin.RouterGroup) {
	petRepo := pet.NewRepository(db)
	petService := pet.NewService(petRepo)
	petHandler := pet.NewHandler(petService)

	group.GET("/pets", checkJWT(), petHandler.Get)
	group.GET("/pet/:id", petHandler.GetById)
	group.POST("/pet", petHandler.Create)
	group.PUT("/pet/:id", petHandler.Update)
	group.DELETE("/pet/:id", petHandler.Delete)
}
