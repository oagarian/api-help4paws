package route

import (
	ctx "context"
	"database/sql"
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
	switch order {
		case "1":
			data, err := database.GetAssociateds(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
			}
			return context.JSON(http.StatusOK, data)
		case "2":
			data, err := database.GetAssociatedsFromLocation(ctx.Background(), db.GetAssociatedsFromLocationParams{Column1: sql.NullString{String: locate, Valid: true}, Limit: int32(intValue)})
			if err != nil {
				log.Println(err)
			}
			return context.JSON(http.StatusOK, data)
		case "3":
			data, err := database.GetAssociatedsInverted(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
			}
			return context.JSON(http.StatusOK, data)
		default: 
			data, err := database.GetAssociateds(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
			}
			return context.JSON(http.StatusOK, data)
	}
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


