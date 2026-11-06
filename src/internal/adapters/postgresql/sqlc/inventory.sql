-- name: CreateInventory :exec
INSERT INTO inventory (profile_id, item_id)
VALUES ($1, $2);

-- name: GetInventory :many
SELECT it.id, it.name, it.value, inv.purchased_at FROM item it INNER JOIN inventory inv ON inv.profile_id = $1 AND inv.item_id = it.id; -- no pagination yet

-- name: GetInventoryItem :one
SELECT purchased_at FROM inventory WHERE profile_id = $1 AND item_id = $2;