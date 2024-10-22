-- name: AdminGetProducts :many
select *
from articulo
;

-- name: AdminGetProductById :one
select *
from articulo
where id = ?
;

-- name: AdminGetProductByCode :one
select *
from articulo
where codigo = ?
;

-- name: CreateProduct :exec
insert into articulo (
id,
codigo,
grupo,
subgrupo,
nombre,
referencia,
marca,
unidad,
existencia,
precio1,
precio2,
precio3,
precio4,
precio5,
precio6,
precio7,
discont,
vta_max,
vta_min,
dctotope,
enpreventa,
comprometido,
vta_minenx,
vta_solofac,
vta_solone,
codbarras,
detalles,
cantbulto,
costo_prom,
util1,
util2,
util3,
fchultcomp,
qtyultcomp,
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

-- name: UpdateProduct :exec
update articulo
set grupo = ?,
    subgrupo = ?,
    nombre = ?,
    referencia = ?,
    marca = ?,
    unidad = ?,
    existencia = ?,
    precio1 = ?,
    precio2 = ?,
    precio3 = ?,
    precio4 = ?,
    precio5 = ?,
    precio6 = ?,
    precio7 = ?,
    discont = ?,
    vta_max = ?,
    vta_min = ?,
    dctotope = ?,
    enpreventa = ?,
    comprometido = ?,
    vta_minenx = ?,
    vta_solofac = ?,
    vta_solone = ?,
    codbarras = ?,
    detalles = ?,
    cantbulto = ?,
    costo_prom = ?,
    util1 = ?,
    util2 = ?,
    util3 = ?,
    fchultcomp = ?,
    qtyultcomp = ?,
    updated_at = NOW()
where id = ?;

-- name: SoftDeleteProduct :exec
update articulo
set deleted_at = NOW()
where id = ?;

-- name: GetAllProducts :many
select *
from articulo
where deleted_at is null
;

-- name: GetProductById :one
select *
from articulo
where id = ? and deleted_at is null
;

-- name: GetProductByCode :one
select *
from articulo
where codigo = ? and deleted_at is null
;

