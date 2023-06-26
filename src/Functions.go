package main

import (
	"database/sql"
	"log"
	"modules_API/internal/db"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func ConnectDatabase() (*db.Queries) {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err)
		return nil
	}

	DatabaseURL := os.Getenv("DATABASE_URL")
	dbconn, err := sql.Open("postgres", DatabaseURL)
	database := db.New(dbconn)
	return database

}


