package user

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
