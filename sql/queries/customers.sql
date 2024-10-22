-- name: AdminGetCustomers :many
select *
from cliente
;

-- name: AdminGetCustomerById :one
select *
from cliente
where id = ?
;

-- name: CreateCustomer :exec
INSERT INTO cliente (
    id,
    codigo,
    nombre,
    direccion,
    telefonos,
    perscont,
    vendedor,
    contribespecial,
    status,
    sector,
    subcodigo,
    precio,
    email,
    kne_activa,
    kne_mtomin,
    noemifac,
    noeminota,
    fchultvta,
    mtoultvta,
    prcdpagdia,
    promdiasp,
    riesgocrd,
    cantdocs,
    totmtodocs,
    prommtodoc,
    diasultvta,
    promdiasvta,
    limcred,
    dolarflete,
    nodolarflete,
    created_at,
    updated_at
)
VALUES (
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
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    NOW(),
    NOW()
);

-- name: UpdateCustomer :exec
UPDATE cliente
SET codigo = ?,
    nombre = ?,
    direccion = ?,
    telefonos = ?,
    perscont = ?,
    vendedor = ?,
    contribespecial = ?,
    status = ?,
    sector = ?,
    subcodigo = ?,
    precio = ?,
    email = ?,
    kne_activa = ?,
    kne_mtomin = ?,
    noemifac = ?,
    noeminota = ?,
    fchultvta = ?,
    mtoultvta = ?,
    prcdpagdia = ?,
    promdiasp = ?,
    riesgocrd = ?,
    cantdocs = ?,
    totmtodocs = ?,
    prommtodoc = ?,
    diasultvta = ?,
    promdiasvta = ?,
    limcred = ?,
    dolarflete = ?,
    nodolarflete = ?,
    updated_at = NOW()
WHERE id = ?;

-- name: SoftDeleteCustomer :exec
UPDATE cliente
SET deleted_at = NOW()
WHERE id = ?;

-- name: GetCustomersByManager :many
select *
from cliente
inner join vendedor on cliente.vendedor = vendedor.codigo
where
    vendedor.supervpor in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetCustomersBySalesman :many
select *
from cliente
where vendedor = ? and deleted_at is null
;

-- name: GetCustomerById :one
select *
from cliente
where id = ? and deleted_at is null
;

