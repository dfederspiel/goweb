package v3

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"rsi.com/go-training/api/auth"
)

type testableAuthHandler struct{}

func (t testableAuthHandler) CurrentUser(c *gin.Context) {
	panic("implement me")
}

func (t testableAuthHandler) Callback(c *gin.Context) {
	panic("implement me")
}

func (t testableAuthHandler) Logout(c *gin.Context) {
	panic("implement me")
}

func (t testableAuthHandler) RequiresAuth(profile auth.AuthProfile) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewTestableAuthHandler() auth.Handler {
	return &testableAuthHandler{}
}

var (
	router *gin.Engine
	rr     *httptest.ResponseRecorder
)

func init() {
	router = gin.Default()
	authHandler := NewTestableAuthHandler()

	a := NewApi(router, authHandler)
	a.ConfigurePetRoutes(NewTestablePetRepository())
	a.ConfigureUserRoutes(NewTestableUserRepository())

	rr = httptest.NewRecorder()
}
