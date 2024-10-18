-- name: AdminGetConfigurations :many
select *
from ke_wcnf_conf
;

-- name: AdminGetConfigurationById :one
select *
from ke_wcnf_conf
where id = ?
;

-- name: InsertConfiguration :exec
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
    username,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateConfiguration :exec
UPDATE ke_wcnf_conf
SET cnfg_idconfig = ?,
    cnfg_clase = ?,
    cnfg_tipo = ?,
    cnfg_valnum = ?,
    cnfg_valsino = ?,
    cnfg_valtxt = ?,
    cnfg_lentxt = ?,
    cnfg_valfch = ?,
    cnfg_activa = ?,
    cnfg_etiq = ?,
    cnfg_ttip = ?,
    username = ?,
    updated_at = NOW() 
WHERE id = ?;

-- name: SoftDeleteConfiguration :exec
UPDATE ke_wcnf_conf
SET deleted_at = NOW()
WHERE id = ?;

-- name: GetConfigurationsByUser :many
select *
from ke_wcnf_conf
where username = ? and deleted_at is null
;

-- name: GetUserConfigurationById :one
select *
from ke_wcnf_conf
where id = ? and username = ? and deleted_at is null
;

