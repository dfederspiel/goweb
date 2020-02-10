package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var keySet *jwk.Set
var store = sessions.NewCookieStore([]byte{97, 17, 116, 200, 253, 170, 104, 186, 89, 180, 92, 37, 4, 157, 157, 90, 18, 73, 219, 37, 49, 127, 85, 179, 67, 65, 72, 254, 101, 126, 218, 58})

type Service interface {
	Get(c *gin.Context) (User, error)
	RegisterOauthCallbackRoute(engine *gin.Engine) gin.IRoutes
	RequiresAuth(profile AuthProfile) gin.HandlerFunc
}

type service struct {
	repo Repository
}

func (s service) Get(c *gin.Context) (User, error) {
	session, _ := store.Get(c.Request, "session")
	email := session.Values["email"]
	return User{
		ID:    "",
		Email: fmt.Sprintf("%v", email),
		Role:  0,
	}, nil
}

func (s service) GetByEmail(email string) (User, error) {
	return s.repo.GetByEmail(email)
}

func (s service) RegisterOauthCallbackRoute(engine *gin.Engine) gin.IRoutes {
	return engine.GET("/callback", func(context *gin.Context) {
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
			"code":          {context.Query("code")},
			"client_id":     {s.ClientId},
			"client_secret": {s.ClientSecret},
			"redirect_uri":  {"http://localhost:8080/callback"},
			"grant_type":    {"authorization_code"},
		}
		response, _ := http.PostForm(s.TokenUri, formData)
		var authResponse AuthResponse
		getJson(response, &authResponse)
		http.SetCookie(context.Writer, &http.Cookie{
			Name:     "token",
			Value:    authResponse.IdToken,
			Expires:  time.Now().Add(120 * time.Minute),
			HttpOnly: true,
		})

		context.Redirect(http.StatusMovedPermanently, "/")
	})
}

func (s service) RequiresAuth(profile AuthProfile) gin.HandlerFunc {
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

		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "token not found")
			return
		}

		claims := &Claims{}
		_, err = jwt.ParseWithClaims(token, claims, validate)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, err.Error())
		}

		user, err := s.GetByEmail(claims.Email)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, "user not found")
		}

		if user.Role > profile.RoleRequired {
			respondWithError(c, http.StatusUnauthorized, "user does not have privileges to perform this action")
		}

		session, _ := store.Get(c.Request, "session")
		session.Values["email"] = user.Email
		session.Save(c.Request, c.Writer)

		c.Next()
	}
}

func NewService(r Repository) Service {
	return &service{r}
}

func getToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			respondWithError(c, http.StatusUnauthorized, err.Error())
		}
		token = auth[1]
	}
	return token, nil
}

func validate(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	}

	keySet, _ = jwk.FetchHTTP("https://www.googleapis.com/oauth2/v3/certs")
	key, _ := getKey(t)
	return key, nil
}

func getKey(token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}
	if key := keySet.LookupKeyID(keyID); len(key) == 1 {
		return key[0].Materialize()
	}
	return nil, errors.New("unable to find key")
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
