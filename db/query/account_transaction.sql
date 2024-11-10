-- name: CreateTransaction :one
INSERT INTO account_transactions (
  account_id,
  amount
) VALUES ( $1, $2 )
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM account_transactions
WHERE id = $1
LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM account_transactions
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;
