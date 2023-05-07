-- name: ListUsers :many
SELECT * from users
ORDER BY id
OFFSET $1
LIMIT $2;

-- name: FindUser :one
SELECT * from users
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * from users
WHERE email = $1;

-- name: CreateAUser :one
INSERT INTO users (
  email
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1;
-- name: DeleteUserByPhone :exec
DELETE FROM users
WHERE phone = $1;

-- name: UpdateUser :exec
UPDATE users SET
  email = $2,
  phone = $3,
  address = $4
WHERE
  id = $1;