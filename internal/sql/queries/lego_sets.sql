-- name: CreateLegoSet :one
INSERT INTO lego_sets (id, serial_number, name, price, theme, year, total_parts, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetLegoSetBySerial :one
SELECT * FROM lego_sets WHERE serial_number = $1;

-- name: ListLegoSets :many
SELECT * FROM lego_sets ORDER BY year DESC;

-- name: UpdateLegoSet :exec
UPDATE lego_sets
SET name = $2, price = $3, theme = $4, year = $5, total_parts = $6
WHERE serial_number = $1;

-- name: DeleteLegoSet :exec
DELETE FROM lego_sets
WHERE serial_number = $1;