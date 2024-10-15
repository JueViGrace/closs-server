-- name: FindAllDocuments :many
select *
from ke_doccti
;

-- name: FindAllDocumentsByCode :many
select *
from ke_doccti
where ke_doccti.vendedor = $1
;

-- name: FindAllDocumentsWithLinesByCode :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.documento = ke_doclmv.documento
where vendedor = $1
;

-- name: FindAllDocumentsWithLines :many
select *
from ke_doccti
left join ke_doclmv on ke_doccti.documento = ke_doclmv.documento
;

