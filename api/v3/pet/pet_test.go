package pet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rsi.com/go-training/mocks"
	"rsi.com/go-training/models"
	"testing"
)

func TestPetService(t *testing.T) {

	r := NewRouter(gin.Default(), NewHandler(NewService(mocks.NewTestablePetRepository())), mocks.NewTestableAuthHandler())
	r.Configure()
	router := r.Engine()
	rr := httptest.NewRecorder()

	t.Run("gets all pets", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/pets", nil)
		router.ServeHTTP(rr, req)

		var p []models.Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, []models.Pet{{
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

		var p models.Pet
		err := json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, models.Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		}, p)
	})

	t.Run("creates a pet", func(t *testing.T) {
		j, err := json.Marshal(models.Pet{
			ID:    "1",
			Name:  "Buddy",
			Age:   4,
			Legs:  true,
			Color: "Green",
		})
		req, _ := http.NewRequest("POST", "/pet", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)

		var p models.Pet
		err = json.NewDecoder(rr.Body).Decode(&p)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, models.Pet{
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
		j, _ := json.Marshal(models.Pet{
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
