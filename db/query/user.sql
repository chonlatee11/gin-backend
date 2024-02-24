-- name: CreateUser :one
INSERT INTO users (email, 
password) 
VALUES ($1, $2) 
RETURNING id, email, create_at, update_at;

-- name: GetUserByEmail :one
SELECT id, email, password, create_at, update_at FROM users WHERE email = $1;