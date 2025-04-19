-- name: CreatePart :one
INSERT INTO lego_parts (serial_number, quantity, name, color, price)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPartByNumber :one
SELECT * FROM lego_parts
WHERE part_number = $1;

-- name: ListParts :many
SELECT * FROM lego_parts
ORDER BY name;

-- name: UpdatePartQuantity :exec
UPDATE lego_parts
SET quantity = $2
WHERE part_number = $1;

-- name: UpdatePartPrice :exec
UPDATE lego_parts
SET price = $2
WHERE part_number = $1;

-- name: DeletePart :exec
DELETE FROM lego_parts
WHERE part_number = $1;