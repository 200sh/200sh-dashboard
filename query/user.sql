-- name: FindUserByProviderID :one
SELECT *
FROM user
WHERE provider_id = ?
LIMIT 1;

-- name: CreateUser :one
INSERT INTO user (provider_id, provider, name, email, status)
VALUES (?, ?, ?, ?, ?)
RETURNING *;