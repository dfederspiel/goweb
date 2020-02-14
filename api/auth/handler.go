package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"rsi.com/go-training/models"
	"time"
)

type Handler interface {
	CurrentUser(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
	RequiresAuth(role models.Role) gin.HandlerFunc
}

type handler struct {
	service Service
}

func (h handler) CurrentUser(c *gin.Context) {
	user, _ := h.service.CurrentUser(c)
	c.JSON(http.StatusOK, user)
}

func (h handler) RequiresAuth(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)

		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.Status(http.StatusUnauthorized)
				respondWithError(c, http.StatusUnauthorized, err.Error())
				return
			}
			// For any other type of error, return a bad request status
			respondWithError(c, http.StatusUnauthorized, err.Error())
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := h.service.GetUserFromToken(token)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, err.Error())
		}

		if user.Role > role {
			respondWithError(c, http.StatusUnauthorized, "user does not have privileges to perform this action")
		}

		c.Next()
	}
}

func (h handler) Callback(c *gin.Context) {
	settings, err := ioutil.ReadFile(os.Getenv("OIDC"))
	if err != nil {
		panic("add oidc.json file from google credentials")
	}

	var oidcSettings models.OIDCSettings
	err = json.Unmarshal(settings, &oidcSettings)
	if err != nil {
		panic("error parsing oidc.json file")
	}

	s := &oidcSettings.WebSettings
	formData := url.Values{
		"code":          {c.Query("code")},
		"client_id":     {s.ClientId},
		"client_secret": {s.ClientSecret},
		"redirect_uri":  {"http://localhost:8080/callback"},
		"grant_type":    {"authorization_code"},
	}
	response, _ := http.PostForm(s.TokenUri, formData)
	var authResponse models.AuthResponse
	json.NewDecoder(response.Body).Decode(&authResponse)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    authResponse.IdToken,
		Expires:  time.Now().Add(120 * time.Minute),
		HttpOnly: true,
	})

	c.Redirect(http.StatusMovedPermanently, "/")
}

func (h handler) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
	c.Redirect(http.StatusPermanentRedirect, "/")
}

func NewHandler(service Service) Handler {
	return &handler{service}
}
