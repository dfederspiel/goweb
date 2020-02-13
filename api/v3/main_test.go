package v3

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/api/auth"
	"rsi.com/go-training/api/v3/pet"
	"rsi.com/go-training/api/v3/user"
	"testing"
)

type testablePetRepository struct{}

func (t testablePetRepository) GetAll() (pets []*pet.Pet, err error) {
	return []*pet.Pet{{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}}, nil
}

func (t testablePetRepository) GetById(id string) (p *pet.Pet, err error) {
	return &pet.Pet{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}, nil
}

func (t testablePetRepository) Create(pet *pet.Pet) (err error) {
	return nil
}

func (t testablePetRepository) Update(pet *pet.Pet) (err error) {
	return nil
}

func (t testablePetRepository) DeleteById(id string) (err error) {
	return nil
}

func NewTestablePetRepository() pet.Repository {
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

func (t testableAuthHandler) RequiresAuth(profile auth.AuthProfile) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewTestableAuthHandler() auth.Handler {
	return &testableAuthHandler{}
}

type testableUserRepository struct{}

func (t testableUserRepository) GetByEmail(email string) (user.User, error) {
	panic("implement me")
}

func NewTestableUserRepository() user.Repository {
	return &testableUserRepository{}
}

func TestPetService(t *testing.T) {

	router := setupRouter()
	rr := httptest.NewRecorder()

	t.Run("get all pets", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v3/pets", nil)
		router.ServeHTTP(rr, req)

		var p []pet.Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, []pet.Pet{{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}}, p)
	})

	t.Run("get pet by id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v3/pet/1", nil)
		router.ServeHTTP(rr, req)

		var p pet.Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, pet.Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}, p)
	})

	t.Run("create pet", func(t *testing.T) {
		j, err := json.Marshal(pet.Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		})
		req, _ := http.NewRequest("POST", "/api/v3/pet", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		var p pet.Pet
		err = json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, pet.Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}, p)
	})

	t.Run("update pet", func(t *testing.T) {
		j, _ := json.Marshal(pet.Pet{
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		})
		req, _ := http.NewRequest("PUT", "/api/v3/pet/1", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})

	t.Run("delete pet", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/api/v3/pet/1", nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	auth := NewTestableAuthHandler()
	a := NewApi(r, auth)
	a.ConfigurePetRoutes(NewTestablePetRepository())
	a.ConfigureUserRoutes(NewTestableUserRepository())
	return r
}
