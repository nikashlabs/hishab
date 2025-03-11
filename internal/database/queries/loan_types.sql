-- name: CreateLoanType :one
INSERT INTO loan_types (user_id, name) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetLoanTypeByID :one
SELECT * FROM loan_types WHERE id = $1;

-- name: GetLoanTypesByUserID :many
SELECT * FROM loan_types WHERE user_id = $1;

-- name: UpdateLoanType :exec
UPDATE loan_types
SET name = $1, modified_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteLoanType :exec
DELETE FROM loan_types WHERE id = $1;
