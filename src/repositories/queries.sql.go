package db

import (
	"context"
	"database/sql"
)

const deleteAssociated = `-- name: DeleteAssociated :exec
DELETE FROM associateds WHERE id = $1
`

func (q *Queries) DeleteAssociated(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAssociated, id)
	return err
}

const getAssociateds = `-- name: GetAssociateds :many
SELECT id, asscname, logoimage, asscdescription, email, contactnumber, pix, street, descriptionaddr FROM associateds ORDER BY id LIMIT $1
`

func (q *Queries) GetAssociateds(ctx context.Context, limit int32) ([]Associated, error) {
	rows, err := q.db.QueryContext(ctx, getAssociateds, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Associated
	for rows.Next() {
		var i Associated
		if err := rows.Scan(
			&i.ID,
			&i.Asscname,
			&i.Logoimage,
			&i.Asscdescription,
			&i.Email,
			&i.Contactnumber,
			&i.Pix,
			&i.Street,
			&i.Descriptionaddr,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssociatedsFromLocation = `-- name: GetAssociatedsFromLocation :many
SELECT id, asscname, logoimage, asscdescription, email, contactnumber, pix, street, descriptionaddr FROM associateds ORDER BY 
CASE 
    WHEN descriptionAddr LIKE '%' || $1 || '%' THEN 0 
    ELSE 1 
END LIMIT $2
`

type GetAssociatedsFromLocationParams struct {
	Column1 sql.NullString
	Limit   int32
}

func (q *Queries) GetAssociatedsFromLocation(ctx context.Context, arg GetAssociatedsFromLocationParams) ([]Associated, error) {
	rows, err := q.db.QueryContext(ctx, getAssociatedsFromLocation, arg.Column1, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Associated
	for rows.Next() {
		var i Associated
		if err := rows.Scan(
			&i.ID,
			&i.Asscname,
			&i.Logoimage,
			&i.Asscdescription,
			&i.Email,
			&i.Contactnumber,
			&i.Pix,
			&i.Street,
			&i.Descriptionaddr,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssociatedsInverted = `-- name: GetAssociatedsInverted :many
SELECT id, asscname, logoimage, asscdescription, email, contactnumber, pix, street, descriptionaddr FROM associateds ORDER BY id DESC LIMIT $1
`

func (q *Queries) GetAssociatedsInverted(ctx context.Context, limit int32) ([]Associated, error) {
	rows, err := q.db.QueryContext(ctx, getAssociatedsInverted, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Associated
	for rows.Next() {
		var i Associated
		if err := rows.Scan(
			&i.ID,
			&i.Asscname,
			&i.Logoimage,
			&i.Asscdescription,
			&i.Email,
			&i.Contactnumber,
			&i.Pix,
			&i.Street,
			&i.Descriptionaddr,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAssociated = `-- name: InsertAssociated :exec
INSERT INTO associateds (asscName, logoImage, asscDescription, email, contactNumber, pix, street, descriptionAddr) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type InsertAssociatedParams struct {
	Asscname        string
	Logoimage       string
	Asscdescription string
	Email           string
	Contactnumber   string
	Pix             string
	Street          string
	Descriptionaddr string
}

func (q *Queries) InsertAssociated(ctx context.Context, arg InsertAssociatedParams) error {
	_, err := q.db.ExecContext(ctx, insertAssociated,
		arg.Asscname,
		arg.Logoimage,
		arg.Asscdescription,
		arg.Email,
		arg.Contactnumber,
		arg.Pix,
		arg.Street,
		arg.Descriptionaddr,
	)
	return err
}

const registerError = `-- name: RegisterError :exec
INSERT INTO logs (timedate, timehour, descriptionerror, wherehappened, solved) VALUES ($1, $2, $3, $4, $5)
`

type RegisterErrorParams struct {
	Timedate         sql.NullTime
	Timehour         sql.NullTime
	Descriptionerror sql.NullString
	Wherehappened    sql.NullString
	Solved           sql.NullBool
}

func (q *Queries) RegisterError(ctx context.Context, arg RegisterErrorParams) error {
	_, err := q.db.ExecContext(ctx, registerError,
		arg.Timedate,
		arg.Timehour,
		arg.Descriptionerror,
		arg.Wherehappened,
		arg.Solved,
	)
	return err
}

const updateAssociated = `-- name: UpdateAssociated :exec
UPDATE associateds SET  asscName = $1, logoImage = $2, asscDescription = $3, email = $4, contactNumber = $5, pix = $6, street = $7, descriptionAddr = $8 WHERE id = $9
`

type UpdateAssociatedParams struct {
	Asscname        string
	Logoimage       string
	Asscdescription string
	Email           string
	Contactnumber   string
	Pix             string
	Street          string
	Descriptionaddr string
	ID              int32
}

func (q *Queries) UpdateAssociated(ctx context.Context, arg UpdateAssociatedParams) error {
	_, err := q.db.ExecContext(ctx, updateAssociated,
		arg.Asscname,
		arg.Logoimage,
		arg.Asscdescription,
		arg.Email,
		arg.Contactnumber,
		arg.Pix,
		arg.Street,
		arg.Descriptionaddr,
		arg.ID,
	)
	return err
}
