-- name: CreateParticipant :one
INSERT INTO participants (
    account,
    event
)
values (
    $1, $2
) RETURNING *;


-- name: ListParticipant :many
SELECT * 
FROM participants
WHERE event = $1
LIMIT $2
OFFSET $3;

-- name: DeleteParticipant :exec
DELETE FROM participants WHERE id = $1;