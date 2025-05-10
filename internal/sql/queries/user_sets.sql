-- name: AddUserSet :one
INSERT INTO user_sets (id, user_id, set_id, price)
SELECT $1, $2, ls.id, $4
FROM lego_sets ls
WHERE ls.serial_number = $3
RETURNING *;

-- name: ListUserSets :many
SELECT * FROM user_sets WHERE user_id = $1 ORDER BY owned_at DESC;

-- name: GetUserSetBySerialNumber :one
SELECT us.*
FROM user_sets us
JOIN lego_sets ls ON us.set_id = ls.id
JOIN users u ON us.user_id = u.id
WHERE u.id = $1 AND ls.serial_number = $2;

-- name: DeleteUserSetBySerial :exec
DELETE FROM user_sets
USING users u, lego_sets ls
WHERE user_sets.user_id = u.id
  AND user_sets.set_id = ls.id
  AND u.id = $1
  AND ls.serial_number = $2;
