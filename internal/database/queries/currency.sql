-- name: CreateCurrency :one
INSERT INTO currency (name, conversion_rate) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetCurrencyByID :one
SELECT * FROM currency WHERE id = $1;

-- name: GetCurrencyByName :one
SELECT * FROM currency WHERE name = $1;

-- name: UpdateCurrency :exec
UPDATE currency
SET name = $1, conversion_rate = $2
WHERE id = $3;

-- name: DeleteCurrency :exec
DELETE FROM currency WHERE id = $1;
