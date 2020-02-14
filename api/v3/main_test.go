package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"rsi.com/go-training/mocks"
	"testing"
)

func TestApi(t *testing.T) {
	t.Run("can initialize the api", func(t *testing.T) {
		a := NewApi(gin.Default(), mocks.NewTestableAuthHandler())
		a.ConfigureUserRoutes(mocks.NewTestableUserRepository())
		a.ConfigurePetRoutes(mocks.NewTestablePetRepository())
		assert.NotNil(t, a)
	})
}
