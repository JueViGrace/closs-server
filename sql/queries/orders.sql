-- name: AdminGetOrders :many
select *
from ke_opti
;

-- name: AdminGetOrdersWithLines :many
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
;

-- name: AdminGetOrderById :one
select *
from ke_opti
where id = ?
;

-- name: AdminGetOrderWithLinesById :one
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
where ke_opti.id = ?
;

-- name: AdminGetOrderByCode :one
select *
from ke_opti
where kti_ndoc = ?
;

-- name: AdminGetOrderLinesByCode :one
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
where ke_opti.kti_ndoc = ?
;

-- name: CreateOrder :exec
insert into ke_opti (
id,
kti_ndoc,
kti_tdoc,
kti_codcli,
kti_nombrecli,
kti_codven,
kti_docsol,
kti_condicion,
kti_tipprec,
kti_totneto,
kti_status,
kti_nroped,
kti_fchdoc,
kti_negesp,
ke_pedstatus,
dolarflete,
complemento,
nro_complemento,
created_at,
updated_at
)
values(
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

-- name: CreateOrderLines :exec
insert into ke_opmv (
    pedido_id,
    articulo_id,
    kti_tdoc,
    kti_ndoc,
    kti_tipprec,
    kmv_codart,
    kmv_nombre,
    kmv_cant,
    kmv_artprec,
    kmv_stot,
    kmv_dctolin,
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
    NOW(),
    NOW()
);

-- name: UpdateOrder :exec
update ke_opti
set kti_tdoc = ?,
    kti_codcli = ?,
    kti_nombrecli = ?,
    kti_codven = ?,
    kti_docsol = ?,
    kti_condicion = ?,
    kti_tipprec = ?,
    kti_totneto = ?,
    kti_status = ?,
    kti_nroped = ?,
    kti_fchdoc = ?,
    kti_negesp = ?,
    ke_pedstatus = ?,
    dolarflete = ?,
    complemento = ?,
    nro_complemento = ?,
    updated_at = NOW()
WHERE id = ?;

-- name: UpdateOrderLines :exec
update ke_opmv
set kti_tdoc = ?,
    kti_tipprec = ?,
    kmv_nombre = ?,
    kmv_cant = ?,
    kmv_artprec = ?,
    kmv_stot = ?,
    kmv_dctolin = ?,
    updated_at = NOW()
WHERE pedido_id = ? and articulo_id = ?;

-- name: SoftDeleteOrder :exec
update ke_opmv
set deleted_at = NOW()
where pedido_id = ?;

-- name: SoftDeleteOrderLines :exec
update ke_opmv
set deleted_at = NOW()
where pedido_id = ?;

-- name: GetOrdersByManager :many
select *
from ke_opti
left join vendedor on ke_opti.kti_codven = vendedor.codigo
where
    vendedor.supervpor in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetOrdersBySalesman :many
select *
from ke_opti
where kti_codven = ? and deleted_at is null
;

-- name: GetOrdersWithLinesByManager :many
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
left join vendedor on ke_opti.kti_codven = vendedor.codigo
where
    vendedor.supervpor in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetOrdersWithLinesBySalesman :many
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
where kti_codven = ? and deleted_at is null
;

-- name: GetOrderById :one
select *
from ke_opti
where id = ? and deleted_at is null
;

-- name: GetOrderWithLinesById :one
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
where ke_opti.id = ? and deleted_at is null
;

-- name: GetOrderByCode :one
select *
from ke_opti
where kti_ndoc = ? and deleted_at is null
;

-- name: GetOrderWithLinesByCode :one
select *
from ke_opti
left join ke_opmv on ke_opti.id = ke_opmv.pedido_id
where ke_opti.kti_ndoc = ? and deleted_at is null
;

