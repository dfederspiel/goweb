package pet

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"rsi.com/go-training/models"
	"testing"
)

func TestPetRepositoryIntegration(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../data/goweb.db")
	fmt.Println(err)
	repo := NewRepository(db)

	t.Run("get all pets", func(t *testing.T) {
		pets, _ := repo.GetAll()
		assert.NotNil(t, pets)
	})

	t.Run("get pet by id", func(t *testing.T) {
		pets, _ := repo.GetById("1")
		assert.NotNil(t, pets)
	})

	t.Run("create pet", func(t *testing.T) {
		err := repo.Create(&models.Pet{
			Name:  "Test Pet",
			Age:   1,
			Color: "Blue",
			Legs:  true,
		})
		assert.Nil(t, err)
	})

	t.Run("update pet", func(t *testing.T) {
		err := repo.Update(&models.Pet{
			ID:    "1",
			Name:  "Test Pet",
			Age:   1,
			Color: "Blue",
			Legs:  true,
		})
		assert.Nil(t, err)
	})

	t.Run("delete pet by id", func(t *testing.T) {
		err := repo.DeleteById("1")
		assert.Nil(t, err)
	})
}
