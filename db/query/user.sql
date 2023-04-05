-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM 
  users
ORDER BY 
  first_name
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateUser :exec
UPDATE 
  users 
SET 
  first_name = $1,
  last_name = $2
WHERE 
  id = $3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
