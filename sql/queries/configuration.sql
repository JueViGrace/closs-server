-- name: AdminGetConfigs :many
select *
from ke_wcnf_conf
;

-- name: AdminGetConfigById :one
select *
from ke_wcnf_conf
where id = ?
;

-- name: CreateConfig :exec
INSERT INTO ke_wcnf_conf (
    id,
    cnfg_idconfig,
    cnfg_clase,
    cnfg_tipo,
    cnfg_valnum,
    cnfg_valsino,
    cnfg_valtxt,
    cnfg_lentxt,
    cnfg_valfch,
    cnfg_activa,
    cnfg_etiq,
    cnfg_ttip,
    user_id,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateConfig :exec
UPDATE ke_wcnf_conf
SET cnfg_clase = ?,
    cnfg_tipo = ?,
    cnfg_valnum = ?,
    cnfg_valsino = ?,
    cnfg_valtxt = ?,
    cnfg_lentxt = ?,
    cnfg_valfch = ?,
    cnfg_activa = ?,
    cnfg_etiq = ?,
    cnfg_ttip = ?,
    updated_at = NOW() 
WHERE id = ?;

-- name: SoftDeleteConfig :exec
UPDATE ke_wcnf_conf
SET deleted_at = NOW()
WHERE id = ?;

-- name: GetConfigByUser :many
select *
from ke_wcnf_conf
where user_id = ? and deleted_at is null
;

-- name: GetUserConfigById :one
select *
from ke_wcnf_conf
where id = ? and user_id = ? and deleted_at is null
;

