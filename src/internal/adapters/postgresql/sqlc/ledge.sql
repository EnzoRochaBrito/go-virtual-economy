-- name: CreateLedge :exec
INSERT INTO profile_ledge (profile_id, exchange_id, amount, movement)
VALUES ($1, $2, $3, $4);

-- name: CreateManyLedges :copyfrom 
INSERT INTO profile_ledge (profile_id, exchange_id, amount, movement)
VALUES ($1, $2, $3, $4);

-- name: GetProfileLedgesFromDate :many
SELECT * FROM profile_ledge WHERE profile_id = $1 AND created_at > $2;