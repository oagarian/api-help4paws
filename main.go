package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string `json:"user"`
	Password string `json:"password"`
}

func handler(c echo.Context) error {
	var user User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Username:", user.Name)
	fmt.Println("Password:", user.Password)
	return c.JSON(http.StatusOK, user)
}

func main() {
	app := echo.New()
	app.GET("/", handler)
	app.Start(":8080")
}
