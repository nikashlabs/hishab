-- name: CreateExpenseCategory :one
INSERT INTO expense_categories (user_id, name) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetExpenseCategoryByID :one
SELECT * FROM expense_categories WHERE id = $1;

-- name: GetExpenseCategoriesByUserID :many
SELECT * FROM expense_categories WHERE user_id = $1;

-- name: UpdateExpenseCategory :exec
UPDATE expense_categories
SET name = $1
WHERE id = $2;

-- name: DeleteExpenseCategory :exec
DELETE FROM expense_categories WHERE id = $1;
