-- name: GetCompanies :many
select *
from ke_dataconex
where deleted_at is null
;

-- name: GetCompanyById :one
select *
from ke_dataconex
where ked_codigo = ? and ked_status = 1 and deleted_at is null
;

-- name: GetAllCompanies :many
select *
from ke_dataconex
;

-- name: GetOneCompanyById :one
select *
from ke_dataconex
where ked_codigo = ?
;

-- name: CreateCompany :exec
INSERT INTO ke_dataconex (
        ked_codigo,
        ked_nombre,
        ked_status,
        ked_enlace,
        ked_agen,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateCompany :exec
UPDATE ke_dataconex
SET ked_codigo = ?,
    ked_nombre = ?,
    ked_status = ?, 
    ked_enlace = ?,
    ked_agen = ?,
    updated_at = NOW()
WHERE ked_codigo = ?;

-- name: DeleteCompany :exec
UPDATE ke_dataconex
SET ked_status = 0,
    deleted_at = NOW()
WHERE ked_codigo = ?;

