-- name: CreateProfile :exec
INSERT INTO profile (id, name) 
VALUES ($1, $2);

-- name: GetProfile :one
SELECT name, created_at FROM profile WHERE id = $1;

-- name: DeleteProfile :exec
DELETE FROM profile WHERE id = $1;

-- name: UpdateProfile :exec
UPDATE profile SET name = $1 WHERE id = $2;