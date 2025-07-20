-- name: GetUserByEmail :one
SELECT 
    id, 
    email, 
    password_hash, 
    is_active, 
    is_verified, 
    role,
    created_at,
    last_login
FROM users 
WHERE email = $1
AND is_active = true
LIMIT 1;

-- name: GetUserByID :one
SELECT 
    id, 
    email, 
    is_active, 
    is_verified, 
    role,
    created_at,
    last_login
FROM users 
WHERE id = $1 
AND is_active = true;

-- name: UpdateLastLogin :exec
UPDATE users 
SET 
    last_login = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
    email, 
    password_hash,
    role
) VALUES (
    $1, $2, $3
) RETURNING id, email, created_at;
