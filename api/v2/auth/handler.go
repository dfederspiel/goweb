package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Handler interface {
	CurrentUser(c *gin.Context)
	Callback(c *gin.Context)
}

type handler struct {
	service Service
}

func (h handler) Callback(c *gin.Context) {
	settings, err := ioutil.ReadFile(os.Getenv("OIDC"))
	if err != nil {
		panic("add oidc.json file from google credentials")
	}

	var oidcSettings OIDCSettings
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
	var authResponse AuthResponse
	getJson(response, &authResponse)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    authResponse.IdToken,
		Expires:  time.Now().Add(120 * time.Minute),
		HttpOnly: true,
	})

	c.Redirect(http.StatusMovedPermanently, "/")
}

func (h handler) CurrentUser(c *gin.Context) {
	user, _ := h.service.CurrentUser(c)
	c.JSON(http.StatusOK, user)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}
