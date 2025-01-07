-- name: CreateOrderLine :exec
insert into closs_order_lines (
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
    ?
);

-- name: UpdateOrderLine :exec
update closs_order_lines set kti_tdoc = ?,
    kti_tipprec = ?,
    kmv_nombre = ?,
    kmv_cant = ?,
    kmv_artprec = ?,
    kmv_stot = ?,
    kmv_dctolin = ?,
    updated_at = ?
WHERE kti_ndoc = ? and kmv_codart = ?;

