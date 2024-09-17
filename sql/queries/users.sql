-- name: CreateUser :exec
INSERT INTO users (id, username, password)
VALUES ($1, $2, $3);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUsernameById :exec
UPDATE users SET username = $1 WHERE id = $2;

-- name: UpdatePasswordById :exec
UPDATE users SET password = $1 WHERE id = $2;
