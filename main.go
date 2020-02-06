package main

import (
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// init is automagically called by go. any package can have an init function
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initializeDB(os.Getenv("DATABASE"))
}

func main() {
	startServer()
}
