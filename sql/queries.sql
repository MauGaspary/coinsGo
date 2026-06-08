-- name: GetUserAccount :one
SELECT account_id, balance FROM accounts
WHERE account_id = $1 LIMIT 1;

-- name: GetUserLoginDetails :one
SELECT account_id, password_hash FROM login_details
WHERE account_id = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (account_id, balance) 
VALUES ($1, 0.0)
RETURNING *;

-- name: CreateLoginDetail :one
INSERT INTO login_details (account_id, password_hash) 
VALUES ($1, $2)
RETURNING *;