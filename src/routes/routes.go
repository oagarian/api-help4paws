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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	finalLocate := cases.Title(language.BrazilianPortuguese).String(locate)
	data := service.SelectRouteResult(order, finalLocate, int32(intValue), context)
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
	var associated model.Associated
	payload := associated
	if err := context.Bind(&payload); err != nil {
		context.String(http.StatusBadRequest, "BadRequest")
	}
	queries := context.QueryParams()
	ID := queries.Get("ID")
	convert, err := strconv.Atoi(ID)
	if err != nil {
		util.RecordLog(err)
	}
	IDInt := int32(convert)
	database.UpdateAssociated(ctx.Background(), db.UpdateAssociatedParams{
		Asscname: payload.AsscName, 
		Logoimage: payload.Logoimage, 
		Asscdescription: payload.Asscdescription,
		Email: payload.Email, 
		Contactnumber: payload.Contactnumber, 
		Pix: payload.Pix, 
		Street: payload.Street, 
		Descriptionaddr: payload.DescriptionAddr,
		ID: IDInt,
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


