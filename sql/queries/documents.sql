-- name: AdminGetDocuments :many
select *
from ke_doccti
;

-- name: AdminGetDocumentsWithLines :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.documento = ke_doclmv.documento
;

-- name: AdminGetDocumentById :one
-- name: InsertDocument :exec
-- name: InsertDocumentLines :exec
-- name: UpdateDocument :exec
-- name: UpdateDocumentLines :exec
-- name: SoftDeleteDocument :exec
-- name: SoftDeleteDocumentLines :exec
-- name: GetDocumentsByManager :many
-- name: GetDocumentsBySalesman :many
select *
from ke_doccti
where ke_doccti.vendedor = ?
;

-- name: GetDocumentsWithLinesByManager :many
-- name: GetDocumentsWithLinesBySalesman :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.documento = ke_doclmv.documento
where ke_doccti.vendedor = ?
;

-- name: GetDocumentById :one


