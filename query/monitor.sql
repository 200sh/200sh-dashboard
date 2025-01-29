-- name: CreateMonitor :one
INSERT INTO monitor(user_id, url)
VALUES (?, ?)
RETURNING *;

-- name: GetMonitorsByUserID :many
SELECT *
FROM monitor
WHERE user_id = ?;

-- name: GetMonitorByUserIDAndMonitorID :one
SELECT *
FROM monitor
WHERE user_id = ?
  and id = ?;

-- name: DeleteMonitor :exec
DELETE FROM monitor
WHERE id = ? AND user_id = ?;
