// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: document_lines.sql

package db

import (
	"context"
)

const createDocumentLine = `-- name: CreateDocumentLine :one
insert into closs_document_lines (
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
    codcoord
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
    ?
)
RETURNING agencia, tipodoc, documento, tipodocv, grupo, subgrupo, origen, codigo, codhijo, pid, nombre, cantidad, cntdevuelt, vndcntdevuelt, dvndmtototal, dpreciofin, dpreciounit, dmontoneto, dmontototal, timpueprc, unidevuelt, fechadoc, vendedor, codcoord
`

type CreateDocumentLineParams struct {
	Agencia       string
	Tipodoc       string
	Documento     string
	Tipodocv      string
	Grupo         string
	Subgrupo      string
	Origen        int64
	Codigo        string
	Codhijo       string
	Pid           string
	Nombre        string
	Cantidad      int64
	Cntdevuelt    int64
	Vndcntdevuelt int64
	Dvndmtototal  float64
	Dpreciofin    float64
	Dpreciounit   float64
	Dmontoneto    float64
	Dmontototal   float64
	Timpueprc     float64
	Unidevuelt    int64
	Fechadoc      string
	Vendedor      string
	Codcoord      string
}

func (q *Queries) CreateDocumentLine(ctx context.Context, arg CreateDocumentLineParams) (ClossDocumentLine, error) {
	row := q.db.QueryRowContext(ctx, createDocumentLine,
		arg.Agencia,
		arg.Tipodoc,
		arg.Documento,
		arg.Tipodocv,
		arg.Grupo,
		arg.Subgrupo,
		arg.Origen,
		arg.Codigo,
		arg.Codhijo,
		arg.Pid,
		arg.Nombre,
		arg.Cantidad,
		arg.Cntdevuelt,
		arg.Vndcntdevuelt,
		arg.Dvndmtototal,
		arg.Dpreciofin,
		arg.Dpreciounit,
		arg.Dmontoneto,
		arg.Dmontototal,
		arg.Timpueprc,
		arg.Unidevuelt,
		arg.Fechadoc,
		arg.Vendedor,
		arg.Codcoord,
	)
	var i ClossDocumentLine
	err := row.Scan(
		&i.Agencia,
		&i.Tipodoc,
		&i.Documento,
		&i.Tipodocv,
		&i.Grupo,
		&i.Subgrupo,
		&i.Origen,
		&i.Codigo,
		&i.Codhijo,
		&i.Pid,
		&i.Nombre,
		&i.Cantidad,
		&i.Cntdevuelt,
		&i.Vndcntdevuelt,
		&i.Dvndmtototal,
		&i.Dpreciofin,
		&i.Dpreciounit,
		&i.Dmontoneto,
		&i.Dmontototal,
		&i.Timpueprc,
		&i.Unidevuelt,
		&i.Fechadoc,
		&i.Vendedor,
		&i.Codcoord,
	)
	return i, err
}

const updateDocumentLine = `-- name: UpdateDocumentLine :one
update closs_document_lines set 
    agencia = ?,
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
    codcoord = ?
WHERE documento = ? and codigo = ?
RETURNING agencia, tipodoc, documento, tipodocv, grupo, subgrupo, origen, codigo, codhijo, pid, nombre, cantidad, cntdevuelt, vndcntdevuelt, dvndmtototal, dpreciofin, dpreciounit, dmontoneto, dmontototal, timpueprc, unidevuelt, fechadoc, vendedor, codcoord
`

type UpdateDocumentLineParams struct {
	Agencia       string
	Tipodoc       string
	Tipodocv      string
	Grupo         string
	Subgrupo      string
	Origen        int64
	Codhijo       string
	Pid           string
	Nombre        string
	Cantidad      int64
	Cntdevuelt    int64
	Vndcntdevuelt int64
	Dvndmtototal  float64
	Dpreciofin    float64
	Dpreciounit   float64
	Dmontoneto    float64
	Dmontototal   float64
	Timpueprc     float64
	Unidevuelt    int64
	Fechadoc      string
	Vendedor      string
	Codcoord      string
	Documento     string
	Codigo        string
}

func (q *Queries) UpdateDocumentLine(ctx context.Context, arg UpdateDocumentLineParams) (ClossDocumentLine, error) {
	row := q.db.QueryRowContext(ctx, updateDocumentLine,
		arg.Agencia,
		arg.Tipodoc,
		arg.Tipodocv,
		arg.Grupo,
		arg.Subgrupo,
		arg.Origen,
		arg.Codhijo,
		arg.Pid,
		arg.Nombre,
		arg.Cantidad,
		arg.Cntdevuelt,
		arg.Vndcntdevuelt,
		arg.Dvndmtototal,
		arg.Dpreciofin,
		arg.Dpreciounit,
		arg.Dmontoneto,
		arg.Dmontototal,
		arg.Timpueprc,
		arg.Unidevuelt,
		arg.Fechadoc,
		arg.Vendedor,
		arg.Codcoord,
		arg.Documento,
		arg.Codigo,
	)
	var i ClossDocumentLine
	err := row.Scan(
		&i.Agencia,
		&i.Tipodoc,
		&i.Documento,
		&i.Tipodocv,
		&i.Grupo,
		&i.Subgrupo,
		&i.Origen,
		&i.Codigo,
		&i.Codhijo,
		&i.Pid,
		&i.Nombre,
		&i.Cantidad,
		&i.Cntdevuelt,
		&i.Vndcntdevuelt,
		&i.Dvndmtototal,
		&i.Dpreciofin,
		&i.Dpreciounit,
		&i.Dmontoneto,
		&i.Dmontototal,
		&i.Timpueprc,
		&i.Unidevuelt,
		&i.Fechadoc,
		&i.Vendedor,
		&i.Codcoord,
	)
	return i, err
}
