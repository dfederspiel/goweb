package v3

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"rsi.com/go-training/api/v3/user"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Run("get user by email", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v3/user/david@federnet.com", nil)
		router.ServeHTTP(rr, req)

		var u user.User
		err := json.NewDecoder(rr.Body).Decode(&u)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, user.User{
			ID:    "1",
			Name:  "Buddy",
			Email: "david@federnet.com",
			Role:  0,
		}, u)
	})
}

type testableUserRepository struct{}

func (t testableUserRepository) GetByEmail(email string) (user.User, error) {
	return user.User{
		ID:    "1",
		Name:  "Buddy",
		Email: "david@federnet.com",
		Role:  0,
	}, nil
}

func NewTestableUserRepository() user.Repository {
	return &testableUserRepository{}
}
