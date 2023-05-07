-- name: CreateAUser :one
INSERT INTO users (
  email
) VALUES (
  $1
)
RETURNING *;