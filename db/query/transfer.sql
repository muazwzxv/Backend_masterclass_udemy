-- name: GetTransfer :one
SELECT * from transfers
WHERE 
  id = $1;

-- name: ListTransferByAccountID :many
SELECT * from transfers
WHERE 
  from_account_id = $1 
OR
  to_account_id = $1
LIMIT $2
OFFSET $3;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;




