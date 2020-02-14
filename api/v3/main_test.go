package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"rsi.com/go-training/api/auth"
	"testing"
)

func TestApi(t *testing.T) {
	t.Run("can initialize the api", func(t *testing.T) {
		a := NewApi(gin.Default(), NewTestableAuthHandler())
		assert.NotNil(t, a)
	})
}

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

func (t testableAuthHandler) RequiresAuth(role auth.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewTestableAuthHandler() auth.Handler {
	return &testableAuthHandler{}
}
