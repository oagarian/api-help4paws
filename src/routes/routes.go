package route

import (
	"context"
	"encoding/json"
	"log"
	util "modules_API/src/utils"
	model "modules_API/src/models"
	"net/http"
	"github.com/labstack/echo/v4"

)

func MainRoute(context echo.Context) error {
	return context.String(http.StatusOK, "Welcome!")
}

func GetAssociatedsRoute(c echo.Context) error {
	db := util.ConnectDatabase()
	data, err := db.GetAssociateds(context.Background())
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, data)
}

func AddAssociatedRoute(context echo.Context) error {
	var associated model.Associated
	err := json.NewDecoder(context.Request().Body).Decode(&associated)
	if err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}
	return context.JSON(http.StatusOK, associated)
}