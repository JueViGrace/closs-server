-- name: AdminGetAllCompanies :many
select *
from ke_dataconex
;

-- name: AdminGetCompanyById :one
select *
from ke_dataconex
where id = ?
;

-- name: CreateCompany :exec
INSERT INTO ke_dataconex (
        id,
        ked_codigo,
        ked_nombre,
        ked_status,
        ked_enlace,
        ked_agen,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateCompany :exec
UPDATE ke_dataconex
SET ked_nombre = ?,
    ked_status = ?, 
    ked_enlace = ?,
    ked_agen = ?,
    updated_at = NOW()
WHERE id = ?;

-- name: SoftDeleteCompany :exec
UPDATE ke_dataconex
SET ked_status = 0,
    deleted_at = NOW()
WHERE id = ?;

-- name: GetCompanyByCode :one
select *
from ke_dataconex
where ked_codigo = ? and ked_status = 1 and deleted_at is null
;

