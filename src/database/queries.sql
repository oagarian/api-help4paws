-- name: InsertAssociated :exec 
INSERT INTO associateds (asscName, logoImage, asscDescription, email, contactNumber, pix, street, descriptionAddr) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetAssociateds :many
SELECT * FROM associateds ORDER BY id LIMIT $1;


-- name: GetAssociatedsFromLocation :many
SELECT * FROM associateds ORDER BY 
CASE 
    WHEN descriptionAddr LIKE '%' || $1 || '%' THEN 0 
    ELSE 1 
END LIMIT $2;

-- name: GetAssociatedsInverted :many 
SELECT * FROM associateds ORDER BY id DESC LIMIT $1;

-- name: DeleteAssociated :exec
DELETE FROM associateds WHERE id = $1;

-- name: UpdateAssociated :exec
UPDATE associateds SET  asscName = $1, logoImage = $2, asscDescription = $3, email = $4, contactNumber = $5, pix = $6, street = $7, descriptionAddr = $8 WHERE id = $9;
