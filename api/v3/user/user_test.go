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
	"rsi.com/go-training/mocks"
	"rsi.com/go-training/models"
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
