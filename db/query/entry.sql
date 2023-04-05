-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries 
WHERE id = $1;

-- name: ListEntriesForAccount :many
SELECT * from entries
WHERE 
  account_id = $1 
LIMIT $2
OFFSET $3;

-- name: ListEntries :many
SELECT * from entries
LIMIT $1
OFFSET $2;

