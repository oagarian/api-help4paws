-- name: InsertAssociated :exec 
INSERT INTO associateds_table (logoImage, asscDescription, email, contactNumber, pix, street, descriptionAddr) 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetAssociateds :many
SELECT * FROM associateds_table ORDER BY id;

-- name: GetAssociatedsFromLocation :many
SELECT * FROM associateds_table WHERE descriptionAddr LIKE '%' || $1 || '%' ORDER BY id;

-- name: GetAssociatedsInverted :many 
SELECT * FROM associateds_table ORDER BY id DESC;

-- name: DeleteAssociated :exec
DELETE FROM associateds_table WHERE id = $1;

-- name: UpdateAssociated :exec
UPDATE associateds_table SET logoImage = $1, asscDescription = $2, email = $3, contactNumber = $4, pix = $5, street = $6, descriptionAddr = $7 WHERE id = $8;
