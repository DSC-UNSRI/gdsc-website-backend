-- name: CreateUser :one
INSERT INTO users (name, birthdate) VALUES ($1, $2) RETURNING *;


-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;