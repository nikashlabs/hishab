-- name: CreateInstallment :one
INSERT INTO installments (
    user_id, account_id, loan_id, title, amount, date, scheduled_record_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetInstallmentByID :one
SELECT * FROM installments WHERE id = $1;

-- name: ListInstallmentsByUser :many
SELECT * FROM installments WHERE user_id = $1 ORDER BY date DESC;

-- name: ListInstallmentsByAccount :many
SELECT * FROM installments WHERE account_id = $1 ORDER BY date DESC;

-- name: ListInstallmentsByLoan :many
SELECT * FROM installments WHERE loan_id = $1 ORDER BY date DESC;

-- name: ListInstallmentsByScheduledRecord :many
SELECT * FROM installments WHERE scheduled_record_id = $1 ORDER BY date DESC;

-- name: UpdateInstallment :exec
UPDATE installments
SET
    user_id = $1,
    account_id = $2,
    loan_id = $3,
    title = $4,
    amount = $5,
    date = $6,
    scheduled_record_id = $7
WHERE id = $8;

-- name: DeleteInstallment :exec
DELETE FROM installments WHERE id = $1;

-- name: DeleteInstallmentsByUser :exec
DELETE FROM installments WHERE user_id = $1;

-- name: DeleteInstallmentsByAccount :exec
DELETE FROM installments WHERE account_id = $1;

-- name: DeleteInstallmentsByLoan :exec
DELETE FROM installments WHERE loan_id = $1;

-- name: DeleteInstallmentsByScheduledRecord :exec
DELETE FROM installments WHERE scheduled_record_id = $1;