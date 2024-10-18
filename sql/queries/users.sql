-- name: AdminGetUsers :many
select *
from usuario
;

-- name: AdminGetUserById :many
select *
from usuario
where id = ?
;

-- name: InsertUser :exec
insert into usuario (
    id,
    username,
    password,
    vendedor_id,
    cliente_id,
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
    ?,
    ?,
    NOW(),
    NOW()
);

-- name: UpdateUser :exec
update usuario
set desactivo = ?,
    ult_sinc = ?,
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

