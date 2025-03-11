-- name: CreateAccount :one
INSERT INTO accounts (user_id, name, balance, currency_id) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountsByUserID :many
SELECT * FROM accounts WHERE user_id = $1;

-- name: UpdateAccount :exec
UPDATE accounts
SET name = $1, balance = $2, currency_id = $3
WHERE id = $4;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
