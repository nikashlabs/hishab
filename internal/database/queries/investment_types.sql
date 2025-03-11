-- name: CreateInvestmentType :one
INSERT INTO investment_types (user_id, name) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetInvestmentTypeByID :one
SELECT * FROM investment_types WHERE id = $1;

-- name: GetInvestmentTypesByUserID :many
SELECT * FROM investment_types WHERE user_id = $1;

-- name: UpdateInvestmentType :exec
UPDATE investment_types
SET name = $1, modified_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteInvestmentType :exec
DELETE FROM investment_types WHERE id = $1;
