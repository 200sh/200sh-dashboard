-- name: CreateMonitor :one
INSERT INTO monitor(user_id, url)
VALUES ($1, $2)
RETURNING *;

-- name: GetMonitorsByUserID :many
SELECT *
FROM monitor
WHERE user_id = $1;

-- name: GetMonitorByUserIDAndMonitorID :one
SELECT *
FROM monitor
WHERE user_id = $1
  and id = $2;

-- name: DeleteMonitor :exec
DELETE FROM monitor
WHERE id = $1 AND user_id = $2;
