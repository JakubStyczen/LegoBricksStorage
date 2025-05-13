-- name: AddPartToSetBySerial :one
INSERT INTO lego_set_parts (set_serial, part_serial, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (set_serial, part_serial) DO UPDATE SET quantity = EXCLUDED.quantity
RETURNING *;

-- name: GetPartsOfSetBySerial :many
SELECT lego_set_parts.set_serial, lego_set_parts.part_serial, lego_set_parts.quantity FROM lego_set_parts WHERE set_serial = $1;

-- name: GetAllPartsInAllSets :many
SELECT 
    lsp.set_serial,
    ls.name AS set_name,
    lsp.part_serial,
    lp.name AS part_name,
    lsp.quantity
FROM lego_set_parts lsp
JOIN lego_sets ls ON lsp.set_serial = ls.serial_number
JOIN lego_parts lp ON lsp.part_serial = lp.serial_number
ORDER BY lsp.set_serial, lsp.part_serial;

-- name: RemovePartFromSet :exec
DELETE FROM lego_set_parts
WHERE set_serial = $1 AND part_serial = $2;