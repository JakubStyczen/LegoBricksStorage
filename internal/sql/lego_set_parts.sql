-- name: AddPartToSet :exec
INSERT INTO lego_set_parts (set_id, part_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (set_id, part_id) DO UPDATE SET quantity = EXCLUDED.quantity;

-- name: GetPartsOfSet :many
SELECT part_id, quantity FROM lego_set_parts WHERE set_id = $1;