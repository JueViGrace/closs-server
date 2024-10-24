-- name: AdminGetSalesmen :many
select *
from vendedor
;

-- name: AdminGetSalesmanById :one
select *
from vendedor
where id = ?
;

-- name: AdminGetSalesmaByCode :one
select *
from vendedor
where codigo = ?
;

-- name: CreateSalesman :exec
insert into vendedor (
    id,
    user_id,
    codigo,
    nombre,
    telefono_1,
    telefono_2,
    telefono_movil,
    status,
    supervpor,
    sector,
    subcodigo,
    email,
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
    ?,
    ?,
    ?,
    NOW(),
    NOW()
);

-- name: UpdateSalesman :exec
update vendedor
set codigo = ?,
    nombre = ?,
    telefono_1 = ?,
    telefono_2 = ?,
    telefono_movil = ?,
    status = ?,
    supervpor = ?,
    sector = ?,
    subcodigo = ?,
    email = ?,
    updated_at = NOW()
where id = ?;

-- name: SoftDeleteSalesman :exec
update vendedor
set deleted_at = NOW()
where id = ?;

-- name: GetSalesmenByManager :many
select *
from vendedor
;

-- name: GetSalesmanById :one
select *
from vendedor
where id = ? and deleted_at is null
;

-- name: GetSalesmanByCode :one
select *
from vendedor
where codigo = ? and deleted_at is null
;

