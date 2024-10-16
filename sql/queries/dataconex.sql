-- name: GetCompanies :many
SELECT * FROM ke_dataconex
WHERE deleted_at IS NULL;

-- name: GetCompanyById :one
SELECT * FROM ke_dataconex
WHERE 
ked_codigo = $1 AND
ked_status = 1 AND
deleted_at IS NULL;

-- name: GetAllCompanies :many
SELECT * FROM ke_dataconex;

-- name: GetOneCompanyById :one
SELECT * FROM ke_dataconex
WHERE ked_codigo = $1;

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
VALUES ($1, $2, $3, $4, $5, NOW(), NOW());

-- name: UpdateCompany :exec
UPDATE ke_dataconex
SET ked_codigo = $1,
    ked_nombre = $2,
    ked_status = $3,
    ked_enlace = $4,
    ked_agen = 5$,
    updated_at = NOW()
WHERE ked_codigo = $6;

-- name: DeleteCompany :exec
UPDATE ke_dataconex
SET ked_status = 0,
    deleted_at = NOW()
WHERE ked_codigo = $1;
