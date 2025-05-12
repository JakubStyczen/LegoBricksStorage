-- name: CreatePart :one
INSERT INTO lego_parts (id, serial_number, quantity, name, color, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPartByNumber :one
SELECT * FROM lego_parts WHERE serial_number = $1;

-- name: ListParts :many
SELECT * FROM lego_parts ORDER BY name;

-- name: UpdatePart :exec
UPDATE lego_parts
SET quantity = $2, name = $3, color = $4, price = $5
WHERE serial_number = $1;

-- name: DeletePart :exec
DELETE FROM lego_parts
WHERE serial_number = $1;