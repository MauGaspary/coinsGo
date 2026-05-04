-- name: GetUserLoginDetails :one
SELECT * FROM login_details
WHERE account_id = $1 LIMIT 1;

-- name: GetUserAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;