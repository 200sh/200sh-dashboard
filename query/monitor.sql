-- name: CreateMonitor :one
INSERT INTO monitor(user_id, url)
VALUES (?, ?)
RETURNING *;

-- name: GetMonitorsByUserID :many
SELECT *
FROM monitor
WHERE user_id = ?;

-- name: GetMonitorByUserID :one
SELECT *
FROM monitor
WHERE user_id = ?
  and id = ?;