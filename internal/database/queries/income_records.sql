-- name: CreateIncomeRecord :one
INSERT INTO income_records (
    user_id, account_id, title, category_id, amount, attachment
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetIncomeRecordByID :one
SELECT * FROM income_records WHERE id = $1;

-- name: ListIncomeRecordsByUser :many
SELECT * FROM income_records WHERE user_id = $1 ORDER BY id DESC;

-- name: ListIncomeRecordsByAccount :many
SELECT * FROM income_records WHERE account_id = $1 ORDER BY id DESC;

-- name: ListIncomeRecordsByCategory :many
SELECT * FROM income_records WHERE category_id = $1 ORDER BY id DESC;

-- name: UpdateIncomeRecord :exec
UPDATE income_records
SET
    user_id = $1,
    account_id = $2,
    title = $3,
    category_id = $4,
    amount = $5,
    attachment = $6
WHERE id = $7;

-- name: DeleteIncomeRecord :exec
DELETE FROM income_records WHERE id = $1;

-- name: DeleteIncomeRecordsByUser :exec
DELETE FROM income_records WHERE user_id = $1;

-- name: DeleteIncomeRecordsByAccount :exec
DELETE FROM income_records WHERE account_id = $1;

-- name: DeleteIncomeRecordsByCategory :exec
DELETE FROM income_records WHERE category_id = $1;