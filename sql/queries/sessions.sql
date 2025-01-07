-- name: GetTokenById :one
select *
from closs_session
where user_id = ?
;

-- name: GetTokenByToken :one
select *
from closs_session
where token = ?
;

-- name: CreateSession :one
INSERT  INTO closs_session(
    token,
    user_id
)
VALUES (?, ?)
RETURNING *;

-- name: UpdateSession :one
UPDATE closs_session SET
    token = ?
WHERE user_id = ?
RETURNING *;

-- name: DeleteSession :exec
delete from closs_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from closs_session
where token = ?
;

