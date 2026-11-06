
-- name: CreateExchange :exec
INSERT INTO exchange (id)
VALUES ($1);

-- name: GetExchangeLedges :many
SELECT * FROM profile_ledge WHERE exchange_id = $1 ORDER BY created_at;