package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/api/auth"
	"testing"
)

func TestUserRepositoryIntegration(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../data/goweb.db")
	fmt.Println(err)
	repo := NewRepository(db)
	t.Run("get by email", func(t *testing.T) {
		user, _ := repo.GetByEmail("david@federnet.com")
		fmt.Println(user)
	})
}

func TestUserService(t *testing.T) {
	r := NewRouter(gin.Default(), NewHandler(NewService(NewTestableUserRepository())), NewTestableAuthHandler())
	r.Configure()
	router := r.Engine()
	rr := httptest.NewRecorder()

	t.Run("get user by email", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/user/david@federnet.com", nil)
		router.ServeHTTP(rr, req)

		var u User
		err := json.NewDecoder(rr.Body).Decode(&u)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, User{
			ID:    "1",
			Name:  "Buddy",
			Email: "david@federnet.com",
			Role:  0,
		}, u)
	})
}

type testableUserRepository struct{}

func (t testableUserRepository) GetByEmail(email string) (User, error) {
	return User{
		ID:    "1",
		Name:  "Buddy",
		Email: "david@federnet.com",
		Role:  0,
	}, nil
}

func NewTestableUserRepository() Repository {
	return &testableUserRepository{}
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
