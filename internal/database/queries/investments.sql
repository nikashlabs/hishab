-- name: CreateInvestment :one
INSERT INTO investments (
    user_id, account_id, title, investment_type_id, amount, date, scheduled_record_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetInvestmentByID :one
SELECT * FROM investments WHERE id = $1;

-- name: ListInvestmentsByUser :many
SELECT * FROM investments WHERE user_id = $1 ORDER BY date DESC;

-- name: ListInvestmentsByAccount :many
SELECT * FROM investments WHERE account_id = $1 ORDER BY date DESC;

-- name: ListInvestmentsByType :many
SELECT * FROM investments WHERE investment_type_id = $1 ORDER BY date DESC;

-- name: ListInvestmentsByScheduledRecord :many
SELECT * FROM investments WHERE scheduled_record_id = $1 ORDER BY date DESC;

-- name: UpdateInvestment :exec
UPDATE investments
SET
    user_id = $1,
    account_id = $2,
    title = $3,
    investment_type_id = $4,
    amount = $5,
    date = $6,
    scheduled_record_id = $7
WHERE id = $8;

-- name: DeleteInvestment :exec
DELETE FROM investments WHERE id = $1;

-- name: DeleteInvestmentsByUser :exec
DELETE FROM investments WHERE user_id = $1;

-- name: DeleteInvestmentsByAccount :exec
DELETE FROM investments WHERE account_id = $1;

-- name: DeleteInvestmentsByType :exec
DELETE FROM investments WHERE investment_type_id = $1;

-- name: DeleteInvestmentsByScheduledRecord :exec
DELETE FROM investments WHERE scheduled_record_id = $1;