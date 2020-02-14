package pet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/api/auth"

	"testing"
)

func TestPetService(t *testing.T) {

	r := NewRouter(gin.Default(), NewHandler(NewService(NewTestablePetRepository())), NewTestableAuthHandler())
	r.Configure()
	router := r.Engine()
	rr := httptest.NewRecorder()

	t.Run("gets all pets", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/pets", nil)
		router.ServeHTTP(rr, req)

		var p []Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, []Pet{{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}}, p)
	})

	t.Run("gets pet by id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/pet/1", nil)
		router.ServeHTTP(rr, req)

		var p Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}, p)
	})

	t.Run("creates a pet", func(t *testing.T) {
		j, err := json.Marshal(Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		})
		req, _ := http.NewRequest("POST", "/pet", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		var p Pet
		err = json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}, p)
	})

	t.Run("won't create pet with bogus body", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/pet", bytes.NewBuffer([]byte("{}")))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)
		fmt.Println(rr.Body)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("updates a pet", func(t *testing.T) {
		j, _ := json.Marshal(Pet{
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		})
		req, _ := http.NewRequest("PUT", "/pet/1", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})

	t.Run("deletes a pet", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/pet/1", nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})
}

type testablePetRepository struct{}

func (t testablePetRepository) GetAll() (pets []*Pet, err error) {
	return []*Pet{{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}}, nil
}

func (t testablePetRepository) GetById(id string) (p *Pet, err error) {
	return &Pet{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}, nil
}

func (t testablePetRepository) Create(pet *Pet) (err error) {
	return nil
}

func (t testablePetRepository) Update(pet *Pet) (err error) {
	return nil
}

func (t testablePetRepository) DeleteById(id string) (err error) {
	return nil
}

func NewTestablePetRepository() Repository {
	return &testablePetRepository{}
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
