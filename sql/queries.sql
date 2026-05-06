-- name: CreateAccount :one
INSERT INTO accounts (account_id)
VALUES ($1);
RETURNING *;

-- name: CreateLoginDetail :one
INSERT INTO login_details (account_id, password_hash)
VALUES ($1, $2);
RETURNING *;

-- name: GetUserLoginDetails :one
SELECT * FROM login_details
WHERE account_id = $1 LIMIT 1;

-- name: GetUserAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;