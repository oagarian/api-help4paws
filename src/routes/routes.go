package route

import (
	ctx "context"
	"log"
	model "modules_API/src/models"
	db "modules_API/src/repositories"
	util "modules_API/src/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MainRoute(context echo.Context) error {
	return context.String(http.StatusOK, "Welcome!")
}

var database = util.ConnectDatabase()


func GetAssociatedsRoute(context echo.Context) error {
	context.Response().Header().Set("Content-Type", "application/json")
	value := context.Param("amount") 
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
	data, err := database.GetAssociateds(ctx.Background(), int32(intValue))
	if err != nil {
		log.Println(err)
	}
	return context.JSON(http.StatusOK, data)
}

func AddAssociatedRoute(context echo.Context) error {
	context.Response().Header().Set("Content-Type", "application/json")
	var associated model.Associated
	payload := associated
	if err := context.Bind(&payload); err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}
	database.InsertAssociated(ctx.Background(), db.InsertAssociatedParams{
		Asscname: payload.AsscName, 
		Logoimage: payload.Logoimage, 
		Email: payload.Email, 
		Contactnumber: payload.Contactnumber, 
		Pix: payload.Pix, 
		Street: payload.Street, 
		Descriptionaddr: payload.DescriptionAddr,
	})
	return context.JSON(http.StatusOK, payload)
}