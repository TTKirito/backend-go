-- name: CreateAccount :one
INSERT INTO accounts (
     owner,
    position,
    gender,
    dob,
    status
) values (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAccount :one
SELECT * 
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: ListAccount :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts
SET position = $2, gender = $3, dob = $4, status = $5
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;