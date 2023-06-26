package route

import (
	"context"
	"encoding/json"
	"log"
	model "modules_API/src/models"
	util "modules_API/src/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MainRoute(context echo.Context) error {
	return context.String(http.StatusOK, "Welcome!")
}

func GetAssociatedsRoute(c echo.Context) error {
	value := c.Param("amount") 
	if value == "" || value == "null"{
		value = "0"
	}
	intValue, err := strconv.Atoi(value)
	if intValue <= 0 {
		intValue = 1
	}

	if err != nil {
		log.Println(err)
	}
	db := util.ConnectDatabase()
	data, err := db.GetAssociateds(context.Background(), int32(intValue))
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