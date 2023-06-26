package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthMiddleware(username, password string, context echo.Context) (bool, error) {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath);
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

type ConfigRateLimit middleware.RateLimiterConfig
func GetLimiter() ConfigRateLimit {
	return ConfigRateLimit{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 20, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(context echo.Context) (string, error) {
			ip := context.RealIP()
			return ip, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string,	err error) error {
        return context.JSON(http.StatusTooManyRequests, nil)
    	},
	}
}