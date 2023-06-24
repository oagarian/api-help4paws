package main

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
)

func MainRoute(context echo.Context) error {
	var user User
	err := json.NewDecoder(context.Request().Body).Decode(&user)
	if err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}
	return context.JSON(http.StatusOK, user)
}

func GetAssociateds(context echo.Context) error {
	return context.String(http.StatusOK, "WIP")
}

func AddAssociated(context echo.Context) error {
	var associated Associated
	err := json.NewDecoder(context.Request().Body).Decode(&associated)
	if err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}
	return context.JSON(http.StatusOK, associated)
}