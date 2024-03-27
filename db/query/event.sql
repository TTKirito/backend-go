-- name: CreateEvent :one
INSERT INTO events (
    title,
    start_time,
    end_time,
    is_emegency,
    owner,
    note,
    type,
    visit_type,
    meeting
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;


-- name: GetEvent :one
SELECT * 
FROM events
WHERE id = $1
LIMIT 1;

-- name: UpdateEvent :one
UPDATE events
SET  title = $2,
     start_time = $3,
     end_time = $4,
     is_emegency = $5,
     owner = $6,
     note = $7,
     type = $8,
     visit_type = $9,
     meeting = $10
WHERE id = $1
RETURNING *;

-- name: ListEvent :many
SELECT * FROM events
WHERE start_time >= $1
AND end_time < $2
ORDER BY start_time
LIMIT $3
OFFSET $4;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;