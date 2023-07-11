package db

import (
	"database/sql"
)

type Associated struct {
	ID              int32
	Asscname        string
	Logoimage       string
	Asscdescription string
	Email           string
	Contactnumber   string
	Pix             string
	Street          string
	Descriptionaddr string
}

type Log struct {
	ID               int32
	Timedate         sql.NullTime
	Timehour         sql.NullTime
	Descriptionerror sql.NullString
	Wherehappened    sql.NullString
	Solved           sql.NullBool
}
