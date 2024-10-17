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
    codigo,
    username,
    email,
    password,
    nombre,
    telefonos,
    telefono_movil,
    direccion,
    sector,
    subcodigo,
    supervpor,
    status,
    desactivo,
    cierre_sesion,
    comisiones,
    ult_sinc,
    version,
    created_at,
    updated_at
)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateUser :exec
UPDATE usuario
set codigo = ?,
    username = ?,
    email = ?,
    password = ?,
    nombre = ?,
    telefonos = ?,
    telefono_movil = ?,
    direccion = ?,
    sector = ?,
    subcodigo = ?,
    supervpor = ?,
    status = ?,
    desactivo = ?,
    cierre_sesion = ?,
    comisiones = ?,
    ult_sinc = ?,
    version = ?,
    updated_at = NOW()
where id = ?;
-- name: SoftDeleteUser :exec
-- name: GetUsersByManager :many
-- name: GetUserBySalesman :one
-- name: GetUserById :one


