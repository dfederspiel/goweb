package v3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/api/v3/pet"
	"testing"
)

func init() {
	router = gin.Default()
	authHandler := NewTestableAuthHandler()
	a := NewApi(router, authHandler)
	a.ConfigurePetRoutes(NewTestablePetRepository())
	rr = httptest.NewRecorder()
}

func TestPetService(t *testing.T) {

	t.Run("gets all pets", func(t *testing.T) {
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

	t.Run("gets pet by id", func(t *testing.T) {
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

	t.Run("creates a pet", func(t *testing.T) {
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

	t.Run("won't create pet with bogus body", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/api/v3/pet", bytes.NewBuffer([]byte("{}")))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)
		fmt.Println(rr.Body)
		assert.Equal(t, 400, rr.Code)
	})

	t.Run("updates a pet", func(t *testing.T) {
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

	t.Run("deletes a pet", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v3/pet/1", nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
	})
}

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
