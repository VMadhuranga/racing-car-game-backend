-- name: CreateUser :exec
INSERT INTO users (id, username, password)
VALUES ($1, $2, $3);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;