package util

import (
	ctx "context"
	"database/sql"
	"log"
	db "modules_API/src/repositories"
	"path"
	"runtime"
	"time"
)

func RecordLog(errSend error) {
	_, archive, _, ok := runtime.Caller(1)
	if !ok {
		return 
	}

	pathArchive := path.Base(archive)


	time := time.Now()
	errorBytes := []byte(errSend.Error())
	errorString := string(errorBytes)
	database := ConnectDatabase()

	database.RegisterError(ctx.Background(), db.RegisterErrorParams{
		Timedate: sql.NullTime{
			Time: time,
			Valid: true,
		},
		Timehour: sql.NullTime{
			Time: time,
			Valid: true,
		},
		Descriptionerror: sql.NullString{
			String: errorString,
			Valid: true,
		},
		Wherehappened: sql.NullString{
			String: pathArchive,
			Valid: true,
		},
		Solved: sql.NullBool{
			Bool: false, 
			Valid: true,
		},
	})
	log.Println("Error assigned")
	
}