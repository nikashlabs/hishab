-- name: CreateLoan :one
INSERT INTO loans (
    user_id, account_id, title, loan_type_id, amount, date
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetLoanByID :one
SELECT * FROM loans WHERE id = $1;

-- name: ListLoansByUser :many
SELECT * FROM loans WHERE user_id = $1 ORDER BY date DESC;

-- name: ListLoansByAccount :many
SELECT * FROM loans WHERE account_id = $1 ORDER BY date DESC;

-- name: ListLoansByType :many
SELECT * FROM loans WHERE loan_type_id = $1 ORDER BY date DESC;

-- name: UpdateLoan :exec
UPDATE loans
SET
    user_id = $1,
    account_id = $2,
    title = $3,
    loan_type_id = $4,
    amount = $5,
    date = $6
WHERE id = $7;

-- name: DeleteLoan :exec
DELETE FROM loans WHERE id = $1;

-- name: DeleteLoansByUser :exec
DELETE FROM loans WHERE user_id = $1;

-- name: DeleteLoansByAccount :exec
DELETE FROM loans WHERE account_id = $1;

-- name: DeleteLoansByType :exec
DELETE FROM loans WHERE loan_type_id = $1;