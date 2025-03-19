-- name: CreateIncomeCategory :one
INSERT INTO income_categories (user_id, name) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetIncomeCategoryByID :one
SELECT * FROM income_categories WHERE id = $1;

-- name: GetIncomeCategoriesByUserID :many
SELECT * FROM income_categories WHERE user_id = $1;

-- name: UpdateIncomeCategory :exec
UPDATE income_categories
SET name = $1
WHERE id = $2;

-- name: DeleteIncomeCategory :exec
DELETE FROM income_categories WHERE id = $1;
