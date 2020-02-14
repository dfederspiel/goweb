package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
	"net/http"
	"rsi.com/go-training/models"
	"strings"
)

var keySet *jwk.Set

type Service interface {
	CurrentUser(c *gin.Context) (models.User, error)
	GetUserFromToken(token string) (models.User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) CurrentUser(c *gin.Context) (models.User, error) {
	token, _ := getToken(c)
	user, _ := s.GetUserFromToken(token)
	return user, nil
}

func (s service) GetUserFromToken(token string) (models.User, error) {
	if token == "" {
		return models.User{}, errors.New("token missing")
	}

	claims := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, claims, validate)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repo.CurrentUser(claims.Email)
	if err != nil {
		return models.User{}, errors.New("user not found")
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
