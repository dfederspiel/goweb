package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/mocks"
	"rsi.com/go-training/models"
	"testing"
)

func TestUserService(t *testing.T) {
	r := NewRouter(gin.Default(), NewHandler(NewService(mocks.NewTestableUserRepository())), mocks.NewTestableAuthHandler())
	r.Configure()
	router := r.Engine()
	rr := httptest.NewRecorder()

	t.Run("get user by email", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/user/david@federnet.com", nil)
		router.ServeHTTP(rr, req)

		var u models.User
		err := json.NewDecoder(rr.Body).Decode(&u)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, models.User{
			ID:    "1",
			Name:  "Buddy",
			Email: "david@federnet.com",
			Role:  0,
		}, u)
	})
}
