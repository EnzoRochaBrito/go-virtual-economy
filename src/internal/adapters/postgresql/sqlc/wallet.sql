-- name: CreateWalletSnapshot :exec
INSERT INTO wallet_daily_snapshot (profile_id, balance, effective_at)
VALUES ($1, $2, $3);

-- name: GetLastSnapshot :one
SELECT * FROM wallet_daily_snapshot WHERE profile_id = $1 ORDER BY effective_at DESC LIMIT 1;