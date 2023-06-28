package route

import (
	ctx "context"
	"log"
	model "modules_API/src/models"
	db "modules_API/src/repositories"
	service "modules_API/src/services"
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
	queries := context.QueryParams()
	order := queries.Get("order")
	locate := queries.Get("locate")
	value := queries.Get("amount")


	if value == "" || value == "null"{
		value = "0"
	}
	intValue, err := strconv.Atoi(value)
	if intValue <= 0 {
		intValue = 1
	}
	if err != nil {
		log.Println(err)
		util.RecordLog(err)
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
		Asscdescription: payload.Asscdescription,
		Email: payload.Email, 
		Contactnumber: payload.Contactnumber, 
		Pix: payload.Pix, 
		Street: payload.Street, 
		Descriptionaddr: payload.DescriptionAddr,
	})
	return context.JSON(http.StatusCreated, payload)
}


func UpdateAssociatedRoute(context echo.Context) error {
	context.Response().Header().Set("Content-Type", "application/json")
	var associated model.UpdateAssociated
	payload := associated
	if err := context.Bind(&payload); err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}

	database.UpdateAssociated(ctx.Background(), db.UpdateAssociatedParams{
		Asscname: payload.AsscName, 
		Logoimage: payload.Logoimage, 
		Asscdescription: payload.Asscdescription,
		Email: payload.Email, 
		Contactnumber: payload.Contactnumber, 
		Pix: payload.Pix, 
		Street: payload.Street, 
		Descriptionaddr: payload.DescriptionAddr,
		ID: payload.ID,
	})
	return context.JSON(http.StatusAccepted, payload)
}

func DeleteAssociatedRoute(context echo.Context) error {
	context.Response().Header().Set("Content-Type", "application/json")
	valueID := context.FormValue("ID")
	ID, err := strconv.Atoi(valueID)
	if err != nil {
		log.Println(err)
		util.RecordLog(err)
	}

	database.DeleteAssociated(ctx.Background(), int32(ID))
	return context.String(http.StatusAccepted, "Deleted successfully!")
}



