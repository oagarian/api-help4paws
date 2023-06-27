package route

import (
	ctx "context"
	model "modules_API/src/models"
	db "modules_API/src/repositories"
	service "modules_API/src/services"
	util "modules_API/src/utils"
	"net/http"
	"strconv"
	"log"
	"github.com/labstack/echo/v4"
)

func MainRoute(context echo.Context) error {
	return context.String(http.StatusOK, "Welcome!")
}

var database = util.ConnectDatabase()


func GetAssociatedsRoute(context echo.Context) error {
	context.Response().Header().Set("Content-Type", "application/json")
	queries := context.QueryParams()
	order := queries.Get("order")
	locate := queries.Get("locate")
	value := queries.Get("value")

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
	data := service.SelectRouteResult(order, locate, int32(intValue), context)
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


