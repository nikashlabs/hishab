-- name: CreateScheduledRecord :one
INSERT INTO scheduled_records (
    user_id, account_id, title, type, category_id, amount, status, date, recurring_interval, recurring
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetScheduledRecordByID :one
SELECT * FROM scheduled_records WHERE id = $1;

-- name: ListScheduledRecordsByUser :many
SELECT * FROM scheduled_records WHERE user_id = $1 ORDER BY date DESC;

-- name: ListScheduledRecordsByAccount :many
SELECT * FROM scheduled_records WHERE account_id = $1 ORDER BY date DESC;

-- name: ListScheduledRecordsByType :many
SELECT * FROM scheduled_records WHERE type = $1 ORDER BY date DESC;

-- name: ListScheduledRecordsByStatus :many
SELECT * FROM scheduled_records WHERE status = $1 ORDER BY date DESC;

-- name: UpdateScheduledRecord :exec
UPDATE scheduled_records
SET
    user_id = $1,
    account_id = $2,
    title = $3,
    type = $4,
    category_id = $5,
    amount = $6,
    status = $7,
    date = $8,
    recurring_interval = $9,
    recurring = $10
WHERE id = $11;

-- name: DeleteScheduledRecord :exec
DELETE FROM scheduled_records WHERE id = $1;

-- name: DeleteScheduledRecordsByUser :exec
DELETE FROM scheduled_records WHERE user_id = $1;

-- name: DeleteScheduledRecordsByAccount :exec
DELETE FROM scheduled_records WHERE account_id = $1;

-- name: DeleteScheduledRecordsByType :exec
DELETE FROM scheduled_records WHERE type = $1;

-- name: DeleteScheduledRecordsByStatus :exec
DELETE FROM scheduled_records WHERE status = $1;