-- name: AdminGetUsers :many
select *
from usuario
;

-- name: AdminGetUserById :many
select *
from usuario
where id = ?
;

-- name: CreateUser :exec
insert into usuario (
    id,
    username,
    password,
    role,
    desactivo,
    ult_sinc,
    version,
    created_at,
    updated_at
)
values (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    NOW(),
    NOW()
);

-- name: UpdateUser :exec
update usuario
set ult_sinc = ?,
    version = ?,
    updated_at = NOW()
where id = ?;

-- name: SoftDeleteUser :exec
update usuario 
set deleted_at = NOW()
where id = ?;

-- name: GetUserById :one
select *
from usuario
where id = ? and deleted_at is null
;

