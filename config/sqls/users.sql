-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1 
LIMIT 1;

-- name: FindUserById :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email ASC;

-- name: CreateUser :one
INSERT INTO  users (
  email
) VALUES (
  $1 
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;
