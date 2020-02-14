package mocks

import (
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/models"
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

func (t testableAuthHandler) RequiresAuth(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewTestableAuthHandler() auth.Handler {
	return &testableAuthHandler{}
}
