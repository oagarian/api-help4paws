package main

import (
	"crypto/subtle"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(username, password string, context echo.Context) (bool, error) {
	err := godotenv.Load();
	if err != nil {
		log.Println(err)
	}
	usernameRequester :=  os.Getenv("USERNAME")
	passwordRequester := os.Getenv("PASSWORD")
	if subtle.ConstantTimeCompare([]byte(username), []byte(usernameRequester)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(passwordRequester)) == 1 {
		return true, nil
	}
	return false, nil
}