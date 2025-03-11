-- name: CreateUser :one
INSERT INTO users (
    name, email, password, is_verified, photo_url
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    name = $1,
    email = $2,
    password = $3,
    is_verified = $4,
    photo_url = $5,
    modified_at = NOW()
WHERE id = $6;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: MarkUserAsVerified :exec
UPDATE users
SET is_verified = TRUE, modified_at = NOW()
WHERE id = $1;

-- name: UpdateLastActive :exec
UPDATE users
SET last_active = NOW()
WHERE id = $1;