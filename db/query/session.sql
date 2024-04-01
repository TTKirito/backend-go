-- name: CreateSession :one
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    is_blocked,
    client_ip,
    expired_at
) values (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;


-- name: GetSession :one
SELECT * FROM sessions 
WHERE refresh_token = $1 LIMIT 1;