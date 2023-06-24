package main

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(username, password string, context echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte("userRequester")) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte("adminPassword")) == 1 {
		return true, nil
	}
	return false, nil
}