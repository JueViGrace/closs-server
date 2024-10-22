-- name: AdminGetStatistics :many
select *
from ke_estadc01
;

-- name: AdminGetStatisticsById :many
select *
from ke_estadc01
where id = ?
;

-- name: AdminGetStatisticsBySaleman :many
select *
from ke_estadc01
where vendedor = ?
;

-- name: CreateStatistic :exec
insert into ke_estadc01(
    id,
    codcoord,
    nomcoord,
    vendedor,
    nombrevend,
    cntpedidos,
    mtopedidos,
    cntfacturas,
    mtofacturas,
    metavend,
    prcmeta,
    cntclientes,
    clivisit,
    prcvisitas,
    lom_montovtas,
    lom_prcvtas,
    lom_prcvisit,
    rlom_montovtas,
    rlom_prcvtas,
    rlom_prcvisit,
    fecha_estad,
    ppgdol_totneto,
    devdol_totneto,
    defdol_totneto,
    totdolcob,
    cntrecl,
    mtorecl,
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
    NOW(),
    NOW()
);

-- name: UpdateStatistic :exec
update ke_estadc01
set codcoord = ?,
    nomcoord = ?,
    vendedor = ?,
    nombrevend = ?,
    cntpedidos = ?,
    mtopedidos = ?,
    cntfacturas = ?,
    mtofacturas = ?,
    metavend = ?,
    prcmeta = ?,
    cntclientes = ?,
    clivisit = ?,
    prcvisitas = ?,
    lom_montovtas = ?,
    lom_prcvtas = ?,
    lom_prcvisit = ?,
    rlom_montovtas = ?,
    rlom_prcvtas = ?,
    rlom_prcvisit = ?,
    fecha_estad = ?,
    ppgdol_totneto = ?,
    devdol_totneto = ?,
    defdol_totneto = ?,
    totdolcob = ?,
    cntrecl = ?,
    mtorecl = ?,
    updated_at = NOW()
where id = ?;

-- name: SoftDeleteStatistic :exec
update ke_estadc01
set deleted_at = NOW()
where id = ?;

-- name: GetStatisticsByManager :many
select *
from ke_estadc01
where
    codcoord in (select kng_codcoord from ke_nivgcia where kng_codgcia = ?)
    and deleted_at is null
;

-- name: GetStatisticsBySalesman :many
select *
from ke_estadc01
where vendedor = ? and deleted_at is null
;

-- name: GetStatisticsById :one
select *
from ke_estadc01
where id = ? and deleted_at is null
;

