-- name: AdminGetDocuments :many
select *
from ke_doccti
;

-- name: AdminGetDocumentsWithLines :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
;

-- name: AdminGetDocumentById :one
select *
from ke_doccti
where id = ?
;

-- name: AdminGetDocumentWithLinesById :one
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
where ke_doccti.id = ?
;

-- name: AdminGetDocumentByCode :one
select *
from ke_doccti
where documento = ?
;

-- name: AdminGetDocumentWithLinesByCode :one
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
where ke_doccti.documento = ?
;

-- name: InsertDocument :exec
insert into ke_doccti (
    id,
    agencia,
    tipodoc,
    documento,
    tipodocv,
    codcliente,
    nombrecli,
    contribesp,
    ruta_parme,
    tipoprecio,
    emision,
    recepcion,
    vence,
    diascred,
    estatusdoc,
    dtotneto,
    dtotimpuest,
    dtotalfinal,
    dtotpagos,
    dtotdescuen,
    dFlete,
    dtotdev,
    dvndmtototal,
    dretencion,
    dretencioniva,
    vendedor,
    codcoord,
    aceptadev,
    kti_negesp,
    bsiva,
    bsflete,
    bsretencion,
    bsretencioniva,
    tasadoc,
    mtodcto,
    fchvencedcto,
    tienedcto,
    cbsret,
    cdret,
    cbsretiva,
    cdretiva,
    cbsrparme,
    cdrparme,
    cbsretflete,
    cdretflete,
    bsmtoiva,
    bsmtofte,
    retmun_mto,
    dolarflete,
    bsretflete,
    dretflete,
    dretmun_mto,
    retivaoblig,
    edoentrega,
    dmtoiva,
    prcdctoaplic,
    montodctodol,
    montodctobs,
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

-- name: InsertDocumentLines :exec
insert into ke_doclmv (
    doc_id,
    articulo_id,
    agencia,
    tipodoc,
    documento,
    tipodocv,
    grupo,
    subgrupo,
    origen,
    codigo,
    codhijo,
    pid,
    nombre,
    cantidad,
    cntdevuelt,
    vndcntdevuelt,
    dvndmtototal,
    dpreciofin,
    dpreciounit,
    dmontoneto,
    dmontototal,
    timpueprc,
    unidevuelt,
    fechadoc,
    vendedor,
    codcoord,
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

-- name: UpdateDocument :exec
update ke_doccti
set agencia = ?,
    tipodoc = ?,
    tipodocv = ?,
    codcliente = ?,
    nombrecli = ?,
    contribesp = ?,
    ruta_parme = ?,
    tipoprecio = ?,
    emision = ?,
    recepcion = ?,
    vence = ?,
    diascred = ?,
    estatusdoc = ?,
    dtotneto = ?,
    dtotimpuest = ?,
    dtotalfinal = ?,
    dtotpagos = ?,
    dtotdescuen = ?,
    dFlete = ?,
    dtotdev = ?,
    dvndmtototal = ?,
    dretencion = ?,
    dretencioniva = ?,
    vendedor = ?,
    codcoord = ?,
    aceptadev = ?,
    kti_negesp = ?,
    bsiva = ?,
    bsflete = ?,
    bsretencion = ?,
    bsretencioniva = ?,
    tasadoc = ?,
    mtodcto = ?,
    fchvencedcto = ?,
    tienedcto = ?,
    cbsret = ?,
    cdret = ?,
    cbsretiva = ?,
    cdretiva = ?,
    cbsrparme = ?,
    cdrparme = ?,
    cbsretflete = ?,
    cdretflete = ?,
    bsmtoiva = ?,
    bsmtofte = ?,
    retmun_mto = ?,
    dolarflete = ?,
    bsretflete = ?,
    dretflete = ?,
    dretmun_mto = ?,
    retivaoblig = ?,
    edoentrega = ?,
    dmtoiva = ?,
    prcdctoaplic = ?,
    montodctodol = ?,
    montodctobs = ?,
    updated_at = NOW()
WHERE id = ?;

-- name: UpdateDocumentLines :exec
update ke_doclmv
set agencia = ?,
    tipodoc = ?,
    tipodocv = ?,
    grupo = ?,
    subgrupo = ?,
    origen = ?,
    codhijo = ?,
    pid = ?,
    nombre = ?,
    cantidad = ?,
    cntdevuelt = ?,
    vndcntdevuelt = ?,
    dvndmtototal = ?,
    dpreciofin = ?,
    dpreciounit = ?,
    dmontoneto = ?,
    dmontototal = ?,
    timpueprc = ?,
    unidevuelt = ?,
    fechadoc = ?,
    vendedor = ?,
    codcoord = ?,
    updated_at = ?
WHERE doc_id = ?;

-- name: SoftDeleteDocument :exec
update ke_doccti
set deleted_at = NOW()
where id = ?;

-- name: SoftDeleteDocumentLines :exec
update ke_doclmv
set deleted_at = NOW()
where doc_id = ?;

-- name: GetDocumentsByManager :many
select *
from ke_doccti
left join vendedor on vendedor.codigo = ke_doccti.vendedor
where
    ke_doccti.codcoord in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetDocumentsBySalesman :many
select *
from ke_doccti
where vendedor = ? and deleted_at is null
;

-- name: GetDocumentsWithLinesByManager :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
left join vendedor on vendedor.codigo = ke_doccti.vendedor
where
    ke_doccti.codcoord in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetDocumentsWithLinesBySalesman :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
where ke_doccti.vendedor = ?
;

-- name: GetDocumentById :one
select *
from ke_doccti
where id = ? and deleted_at is null
;

-- name: GetDocumentWithLinesById :one
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
where ke_doccti.id = ? and deleted_at is null
;

-- name: GetDocumentByCode :one
select *
from ke_doccti
where ke_doccti.documento = ? and deleted_at is null
;

-- name: GetDocumentWithLinesByCode :one
select *
from ke_doccti
left join ke_doclmv on ke_doccti.id = ke_doclmv.doc_id
where ke_doccti.documento = ? and deleted_at is null
;

