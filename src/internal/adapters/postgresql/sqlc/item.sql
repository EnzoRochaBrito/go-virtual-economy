-- name: CreateItem :exec
INSERT INTO item (id, creator_id, name, value)
VALUES ($1, $2, $3, $4);

-- name: GetItem :one
SELECT * FROM item WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM item WHERE id = $1;
