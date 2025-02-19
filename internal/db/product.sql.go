// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
;

insert into closs_product (
    codigo,
    grupo,
    subgrupo,
    nombre,
    referencia,
    marca,
    unidad,
    existencia,
    precio1,
    precio2,
    precio3,
    precio4,
    precio5,
    precio6,
    precio7,
    discont,
    vta_max,
    vta_min,
    dctotope,
    preventa,
    comprometido,
    vta_minenx,
    vta_solofac,
    vta_solone,
    codbarras,
    detalles,
    cantbulto,
    costo_prom,
    util1,
    util2,
    util3,
    fchultcomp,
    qtyultcomp,
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
    ?
)
RETURNING codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
`

type CreateProductParams struct {
	Codigo       string
	Grupo        string
	Subgrupo     string
	Nombre       string
	Referencia   string
	Marca        string
	Unidad       string
	Existencia   int64
	Precio1      float64
	Precio2      float64
	Precio3      float64
	Precio4      float64
	Precio5      float64
	Precio6      float64
	Precio7      float64
	Discont      int64
	VtaMax       int64
	VtaMin       int64
	Dctotope     float64
	Preventa     int64
	Comprometido int64
	VtaMinenx    int64
	VtaSolofac   int64
	VtaSolone    int64
	Codbarras    int64
	Detalles     string
	Cantbulto    int64
	CostoProm    float64
	Util1        float64
	Util2        float64
	Util3        float64
	Fchultcomp   string
	Qtyultcomp   int64
	CreatedAt    string
	UpdatedAt    string
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (ClossProduct, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Codigo,
		arg.Grupo,
		arg.Subgrupo,
		arg.Nombre,
		arg.Referencia,
		arg.Marca,
		arg.Unidad,
		arg.Existencia,
		arg.Precio1,
		arg.Precio2,
		arg.Precio3,
		arg.Precio4,
		arg.Precio5,
		arg.Precio6,
		arg.Precio7,
		arg.Discont,
		arg.VtaMax,
		arg.VtaMin,
		arg.Dctotope,
		arg.Preventa,
		arg.Comprometido,
		arg.VtaMinenx,
		arg.VtaSolofac,
		arg.VtaSolone,
		arg.Codbarras,
		arg.Detalles,
		arg.Cantbulto,
		arg.CostoProm,
		arg.Util1,
		arg.Util2,
		arg.Util3,
		arg.Fchultcomp,
		arg.Qtyultcomp,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i ClossProduct
	err := row.Scan(
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Discont,
		&i.Existencia,
		&i.VtaMax,
		&i.VtaMin,
		&i.VtaMinenx,
		&i.Comprometido,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Preventa,
		&i.Dctotope,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.Images,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
delete from closs_product
where codigo = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, codigo string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, codigo)
	return err
}

const getExistingProductByCode = `-- name: GetExistingProductByCode :one
;

select codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
from closs_product
where
    codigo = ?
    and (existencia - comprometido) > 0
    and discont = 0
    and deleted_at is null
`

func (q *Queries) GetExistingProductByCode(ctx context.Context, codigo string) (ClossProduct, error) {
	row := q.db.QueryRowContext(ctx, getExistingProductByCode, codigo)
	var i ClossProduct
	err := row.Scan(
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Discont,
		&i.Existencia,
		&i.VtaMax,
		&i.VtaMin,
		&i.VtaMinenx,
		&i.Comprometido,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Preventa,
		&i.Dctotope,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.Images,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getExistingProducts = `-- name: GetExistingProducts :many
;

select codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
from closs_product
where (existencia - comprometido) > 0 and discont = 0 and deleted_at is null
`

func (q *Queries) GetExistingProducts(ctx context.Context) ([]ClossProduct, error) {
	rows, err := q.db.QueryContext(ctx, getExistingProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossProduct
	for rows.Next() {
		var i ClossProduct
		if err := rows.Scan(
			&i.Codigo,
			&i.Grupo,
			&i.Subgrupo,
			&i.Nombre,
			&i.Referencia,
			&i.Marca,
			&i.Unidad,
			&i.Discont,
			&i.Existencia,
			&i.VtaMax,
			&i.VtaMin,
			&i.VtaMinenx,
			&i.Comprometido,
			&i.Precio1,
			&i.Precio2,
			&i.Precio3,
			&i.Precio4,
			&i.Precio5,
			&i.Precio6,
			&i.Precio7,
			&i.Preventa,
			&i.Dctotope,
			&i.VtaSolofac,
			&i.VtaSolone,
			&i.Codbarras,
			&i.Detalles,
			&i.Cantbulto,
			&i.CostoProm,
			&i.Util1,
			&i.Util2,
			&i.Util3,
			&i.Fchultcomp,
			&i.Qtyultcomp,
			&i.Images,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductByCode = `-- name: GetProductByCode :one
;

select codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
from closs_product
where codigo = ?
`

func (q *Queries) GetProductByCode(ctx context.Context, codigo string) (ClossProduct, error) {
	row := q.db.QueryRowContext(ctx, getProductByCode, codigo)
	var i ClossProduct
	err := row.Scan(
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Discont,
		&i.Existencia,
		&i.VtaMax,
		&i.VtaMin,
		&i.VtaMinenx,
		&i.Comprometido,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Preventa,
		&i.Dctotope,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.Images,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
select codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
from closs_product
`

func (q *Queries) GetProducts(ctx context.Context) ([]ClossProduct, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossProduct
	for rows.Next() {
		var i ClossProduct
		if err := rows.Scan(
			&i.Codigo,
			&i.Grupo,
			&i.Subgrupo,
			&i.Nombre,
			&i.Referencia,
			&i.Marca,
			&i.Unidad,
			&i.Discont,
			&i.Existencia,
			&i.VtaMax,
			&i.VtaMin,
			&i.VtaMinenx,
			&i.Comprometido,
			&i.Precio1,
			&i.Precio2,
			&i.Precio3,
			&i.Precio4,
			&i.Precio5,
			&i.Precio6,
			&i.Precio7,
			&i.Preventa,
			&i.Dctotope,
			&i.VtaSolofac,
			&i.VtaSolone,
			&i.Codbarras,
			&i.Detalles,
			&i.Cantbulto,
			&i.CostoProm,
			&i.Util1,
			&i.Util2,
			&i.Util3,
			&i.Fchultcomp,
			&i.Qtyultcomp,
			&i.Images,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const softDeleteProduct = `-- name: SoftDeleteProduct :exec
update closs_product set 
    discont = 1,
    updated_at = ?,
    deleted_at = ?
where codigo = ?
`

type SoftDeleteProductParams struct {
	UpdatedAt string
	DeletedAt sql.NullString
	Codigo    string
}

func (q *Queries) SoftDeleteProduct(ctx context.Context, arg SoftDeleteProductParams) error {
	_, err := q.db.ExecContext(ctx, softDeleteProduct, arg.UpdatedAt, arg.DeletedAt, arg.Codigo)
	return err
}

const updateProduct = `-- name: UpdateProduct :one
update closs_product set 
    grupo = ?,
    subgrupo = ?,
    nombre = ?,
    referencia = ?,
    marca = ?,
    unidad = ?,
    existencia = ?,
    precio1 = ?,
    precio2 = ?,
    precio3 = ?,
    precio4 = ?,
    precio5 = ?,
    precio6 = ?,
    precio7 = ?,
    discont = ?,
    vta_max = ?,
    vta_min = ?,
    dctotope = ?,
    preventa = ?,
    comprometido = ?,
    vta_minenx = ?,
    vta_solofac = ?,
    vta_solone = ?,
    codbarras = ?,
    detalles = ?,
    cantbulto = ?,
    costo_prom = ?,
    util1 = ?,
    util2 = ?,
    util3 = ?,
    fchultcomp = ?,
    qtyultcomp = ?,
    updated_at = ?
where codigo = ?
RETURNING codigo, grupo, subgrupo, nombre, referencia, marca, unidad, discont, existencia, vta_max, vta_min, vta_minenx, comprometido, precio1, precio2, precio3, precio4, precio5, precio6, precio7, preventa, dctotope, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, images, created_at, updated_at, deleted_at
`

type UpdateProductParams struct {
	Grupo        string
	Subgrupo     string
	Nombre       string
	Referencia   string
	Marca        string
	Unidad       string
	Existencia   int64
	Precio1      float64
	Precio2      float64
	Precio3      float64
	Precio4      float64
	Precio5      float64
	Precio6      float64
	Precio7      float64
	Discont      int64
	VtaMax       int64
	VtaMin       int64
	Dctotope     float64
	Preventa     int64
	Comprometido int64
	VtaMinenx    int64
	VtaSolofac   int64
	VtaSolone    int64
	Codbarras    int64
	Detalles     string
	Cantbulto    int64
	CostoProm    float64
	Util1        float64
	Util2        float64
	Util3        float64
	Fchultcomp   string
	Qtyultcomp   int64
	UpdatedAt    string
	Codigo       string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (ClossProduct, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Grupo,
		arg.Subgrupo,
		arg.Nombre,
		arg.Referencia,
		arg.Marca,
		arg.Unidad,
		arg.Existencia,
		arg.Precio1,
		arg.Precio2,
		arg.Precio3,
		arg.Precio4,
		arg.Precio5,
		arg.Precio6,
		arg.Precio7,
		arg.Discont,
		arg.VtaMax,
		arg.VtaMin,
		arg.Dctotope,
		arg.Preventa,
		arg.Comprometido,
		arg.VtaMinenx,
		arg.VtaSolofac,
		arg.VtaSolone,
		arg.Codbarras,
		arg.Detalles,
		arg.Cantbulto,
		arg.CostoProm,
		arg.Util1,
		arg.Util2,
		arg.Util3,
		arg.Fchultcomp,
		arg.Qtyultcomp,
		arg.UpdatedAt,
		arg.Codigo,
	)
	var i ClossProduct
	err := row.Scan(
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Discont,
		&i.Existencia,
		&i.VtaMax,
		&i.VtaMin,
		&i.VtaMinenx,
		&i.Comprometido,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Preventa,
		&i.Dctotope,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.Images,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
