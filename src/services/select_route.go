package service

import (
	ctx "context"
	db "modules_API/src/repositories"
	util "modules_API/src/utils"
	"database/sql"
	"log"
	"github.com/labstack/echo/v4"
)

var database = util.ConnectDatabase()
func SelectRouteResult(order, locate string, intValue int32, context echo.Context) ([]db.Associated) {
	switch order {
		case "1":
			data, err := database.GetAssociateds(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
				util.RecordLog(err)
			}
			return data
		case "2":
			data, err := database.GetAssociatedsFromLocation(ctx.Background(), db.GetAssociatedsFromLocationParams{Column1: sql.NullString{String: locate, Valid: true}, Limit: int32(intValue)})
			if err != nil {
				log.Println(err)
				util.RecordLog(err)
			}
			return data
		case "3":
			data, err := database.GetAssociatedsInverted(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
				util.RecordLog(err)
			}
			return data
		default: 
			data, err := database.GetAssociateds(ctx.Background(), int32(intValue))
			if err != nil {
				log.Println(err)
				util.RecordLog(err)
			}
			return data
	}
}