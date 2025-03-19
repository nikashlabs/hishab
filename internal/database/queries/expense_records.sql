-- name: CreateExpenseRecord :one
INSERT INTO expense_records (
    user_id, account_id, title, category_id, amount, attachment
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetExpenseRecordByID :one
SELECT * FROM expense_records WHERE id = $1;

-- name: ListExpenseRecordsByUser :many
SELECT * FROM expense_records WHERE user_id = $1 ORDER BY id DESC;

-- name: ListExpenseRecordsByAccount :many
SELECT * FROM expense_records WHERE account_id = $1 ORDER BY id DESC;

-- name: ListExpenseRecordsByCategory :many
SELECT * FROM expense_records WHERE category_id = $1 ORDER BY id DESC;

-- name: UpdateExpenseRecord :exec
UPDATE expense_records
SET
    user_id = $1,
    account_id = $2,
    title = $3,
    category_id = $4,
    amount = $5,
    attachment = $6
WHERE id = $7;

-- name: DeleteExpenseRecord :exec
DELETE FROM expense_records WHERE id = $1;

-- name: DeleteExpenseRecordsByUser :exec
DELETE FROM expense_records WHERE user_id = $1;

-- name: DeleteExpenseRecordsByAccount :exec
DELETE FROM expense_records WHERE account_id = $1;

-- name: DeleteExpenseRecordsByCategory :exec
DELETE FROM expense_records WHERE category_id = $1;