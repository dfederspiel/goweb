package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var keySet *jwk.Set
var store = sessions.NewCookieStore([]byte{97, 17, 116, 200, 253, 170, 104, 186, 89, 180, 92, 37, 4, 157, 157, 90, 18, 73, 219, 37, 49, 127, 85, 179, 67, 65, 72, 254, 101, 126, 218, 58})

type Service interface {
	RequiresAuth(profile AuthProfile) gin.HandlerFunc
	CurrentUser(c *gin.Context) (User, error)
	GetUserFromToken(token string) (User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) CurrentUser(c *gin.Context) (User, error) {
	token, _ := getToken(c)
	user, _ := s.GetUserFromToken(token)
	return user, nil
}

func (s service) GetUserFromToken(token string) (User, error) {
	if token == "" {
		return User{}, errors.New("token missing")
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, validate)
	if err != nil {
		return User{}, err
	}

	user, err := s.repo.CurrentUser(claims.Email)
	if err != nil {
		return User{}, errors.New("user not found")
		//respondWithError(c, http.StatusUnauthorized, "user not found")
	}

	return user, nil
}

func getToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			respondWithError(c, http.StatusUnauthorized, err.Error())
		} else {
			token = auth[1]
		}
	}
	return token, nil
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

		user, err := s.GetUserFromToken(token)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, err.Error())
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
