// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customer.sql

package db

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
;

INSERT INTO closs_customer (
    codigo,
    nombre,
    email,
    direccion,
    telefonos,
    perscont,
    vendedor,
    contribespecial,
    status,
    sector,
    subsector,
    precio,
    kne_activa,
    kne_mtomin,
    noemifac,
    noeminota,
    fchultvta,
    mtoultvta,
    prcdpagdia,
    promdiasp,
    riesgocrd,
    cantdocs,
    totmtodocs,
    prommtodoc,
    diasultvta,
    promdiasvta,
    limcred,
    fchcrea,
    dolarflete,
    nodolarflete,
    created_at,
    updated_at
)
VALUES (
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
RETURNING codigo, nombre, email, direccion, telefonos, perscont, vendedor, contribespecial, status, sector, subsector, precio, kne_activa, kne_mtomin, noemifac, noeminota, fchultvta, mtoultvta, prcdpagdia, promdiasp, riesgocrd, cantdocs, totmtodocs, prommtodoc, diasultvta, promdiasvta, limcred, fchcrea, dolarflete, nodolarflete, created_at, updated_at
`

type CreateCustomerParams struct {
	Codigo          string
	Nombre          string
	Email           string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial int64
	Status          int64
	Sector          string
	Subsector       string
	Precio          int64
	KneActiva       int64
	KneMtomin       float64
	Noemifac        int64
	Noeminota       int64
	Fchultvta       string
	Mtoultvta       float64
	Prcdpagdia      float64
	Promdiasp       float64
	Riesgocrd       float64
	Cantdocs        int64
	Totmtodocs      float64
	Prommtodoc      float64
	Diasultvta      float64
	Promdiasvta     float64
	Limcred         float64
	Fchcrea         string
	Dolarflete      int64
	Nodolarflete    int64
	CreatedAt       string
	UpdatedAt       string
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (ClossCustomer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.Codigo,
		arg.Nombre,
		arg.Email,
		arg.Direccion,
		arg.Telefonos,
		arg.Perscont,
		arg.Vendedor,
		arg.Contribespecial,
		arg.Status,
		arg.Sector,
		arg.Subsector,
		arg.Precio,
		arg.KneActiva,
		arg.KneMtomin,
		arg.Noemifac,
		arg.Noeminota,
		arg.Fchultvta,
		arg.Mtoultvta,
		arg.Prcdpagdia,
		arg.Promdiasp,
		arg.Riesgocrd,
		arg.Cantdocs,
		arg.Totmtodocs,
		arg.Prommtodoc,
		arg.Diasultvta,
		arg.Promdiasvta,
		arg.Limcred,
		arg.Fchcrea,
		arg.Dolarflete,
		arg.Nodolarflete,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i ClossCustomer
	err := row.Scan(
		&i.Codigo,
		&i.Nombre,
		&i.Email,
		&i.Direccion,
		&i.Telefonos,
		&i.Perscont,
		&i.Vendedor,
		&i.Contribespecial,
		&i.Status,
		&i.Sector,
		&i.Subsector,
		&i.Precio,
		&i.KneActiva,
		&i.KneMtomin,
		&i.Noemifac,
		&i.Noeminota,
		&i.Fchultvta,
		&i.Mtoultvta,
		&i.Prcdpagdia,
		&i.Promdiasp,
		&i.Riesgocrd,
		&i.Cantdocs,
		&i.Totmtodocs,
		&i.Prommtodoc,
		&i.Diasultvta,
		&i.Promdiasvta,
		&i.Limcred,
		&i.Fchcrea,
		&i.Dolarflete,
		&i.Nodolarflete,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCustomerByCode = `-- name: GetCustomerByCode :one
;

select codigo, nombre, email, direccion, telefonos, perscont, vendedor, contribespecial, status, sector, subsector, precio, kne_activa, kne_mtomin, noemifac, noeminota, fchultvta, mtoultvta, prcdpagdia, promdiasp, riesgocrd, cantdocs, totmtodocs, prommtodoc, diasultvta, promdiasvta, limcred, fchcrea, dolarflete, nodolarflete, created_at, updated_at
from closs_customer
where codigo = ?
`

func (q *Queries) GetCustomerByCode(ctx context.Context, codigo string) (ClossCustomer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByCode, codigo)
	var i ClossCustomer
	err := row.Scan(
		&i.Codigo,
		&i.Nombre,
		&i.Email,
		&i.Direccion,
		&i.Telefonos,
		&i.Perscont,
		&i.Vendedor,
		&i.Contribespecial,
		&i.Status,
		&i.Sector,
		&i.Subsector,
		&i.Precio,
		&i.KneActiva,
		&i.KneMtomin,
		&i.Noemifac,
		&i.Noeminota,
		&i.Fchultvta,
		&i.Mtoultvta,
		&i.Prcdpagdia,
		&i.Promdiasp,
		&i.Riesgocrd,
		&i.Cantdocs,
		&i.Totmtodocs,
		&i.Prommtodoc,
		&i.Diasultvta,
		&i.Promdiasvta,
		&i.Limcred,
		&i.Fchcrea,
		&i.Dolarflete,
		&i.Nodolarflete,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCustomers = `-- name: GetCustomers :many
select codigo, nombre, email, direccion, telefonos, perscont, vendedor, contribespecial, status, sector, subsector, precio, kne_activa, kne_mtomin, noemifac, noeminota, fchultvta, mtoultvta, prcdpagdia, promdiasp, riesgocrd, cantdocs, totmtodocs, prommtodoc, diasultvta, promdiasvta, limcred, fchcrea, dolarflete, nodolarflete, created_at, updated_at
from closs_customer
`

func (q *Queries) GetCustomers(ctx context.Context) ([]ClossCustomer, error) {
	rows, err := q.db.QueryContext(ctx, getCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossCustomer
	for rows.Next() {
		var i ClossCustomer
		if err := rows.Scan(
			&i.Codigo,
			&i.Nombre,
			&i.Email,
			&i.Direccion,
			&i.Telefonos,
			&i.Perscont,
			&i.Vendedor,
			&i.Contribespecial,
			&i.Status,
			&i.Sector,
			&i.Subsector,
			&i.Precio,
			&i.KneActiva,
			&i.KneMtomin,
			&i.Noemifac,
			&i.Noeminota,
			&i.Fchultvta,
			&i.Mtoultvta,
			&i.Prcdpagdia,
			&i.Promdiasp,
			&i.Riesgocrd,
			&i.Cantdocs,
			&i.Totmtodocs,
			&i.Prommtodoc,
			&i.Diasultvta,
			&i.Promdiasvta,
			&i.Limcred,
			&i.Fchcrea,
			&i.Dolarflete,
			&i.Nodolarflete,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getCustomersByManager = `-- name: GetCustomersByManager :many
;

select closs_customer.codigo, closs_customer.nombre, closs_customer.email, closs_customer.direccion, closs_customer.telefonos, closs_customer.perscont, closs_customer.vendedor, closs_customer.contribespecial, closs_customer.status, closs_customer.sector, closs_customer.subsector, closs_customer.precio, closs_customer.kne_activa, closs_customer.kne_mtomin, closs_customer.noemifac, closs_customer.noeminota, closs_customer.fchultvta, closs_customer.mtoultvta, closs_customer.prcdpagdia, closs_customer.promdiasp, closs_customer.riesgocrd, closs_customer.cantdocs, closs_customer.totmtodocs, closs_customer.prommtodoc, closs_customer.diasultvta, closs_customer.promdiasvta, closs_customer.limcred, closs_customer.fchcrea, closs_customer.dolarflete, closs_customer.nodolarflete, closs_customer.created_at, closs_customer.updated_at
from closs_customer
left join closs_salesman on closs_customer.vendedor = closs_salesman.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where closs_managers.kng_codgcia =?1
order by
    closs_salesman.supervpor asc, closs_customer.vendedor asc, closs_customer.nombre asc
`

func (q *Queries) GetCustomersByManager(ctx context.Context, manager string) ([]ClossCustomer, error) {
	rows, err := q.db.QueryContext(ctx, getCustomersByManager, manager)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossCustomer
	for rows.Next() {
		var i ClossCustomer
		if err := rows.Scan(
			&i.Codigo,
			&i.Nombre,
			&i.Email,
			&i.Direccion,
			&i.Telefonos,
			&i.Perscont,
			&i.Vendedor,
			&i.Contribespecial,
			&i.Status,
			&i.Sector,
			&i.Subsector,
			&i.Precio,
			&i.KneActiva,
			&i.KneMtomin,
			&i.Noemifac,
			&i.Noeminota,
			&i.Fchultvta,
			&i.Mtoultvta,
			&i.Prcdpagdia,
			&i.Promdiasp,
			&i.Riesgocrd,
			&i.Cantdocs,
			&i.Totmtodocs,
			&i.Prommtodoc,
			&i.Diasultvta,
			&i.Promdiasvta,
			&i.Limcred,
			&i.Fchcrea,
			&i.Dolarflete,
			&i.Nodolarflete,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getCustomersBySalesman = `-- name: GetCustomersBySalesman :many
;

select codigo, nombre, email, direccion, telefonos, perscont, vendedor, contribespecial, status, sector, subsector, precio, kne_activa, kne_mtomin, noemifac, noeminota, fchultvta, mtoultvta, prcdpagdia, promdiasp, riesgocrd, cantdocs, totmtodocs, prommtodoc, diasultvta, promdiasvta, limcred, fchcrea, dolarflete, nodolarflete, created_at, updated_at
from closs_customer
where closs_customer.vendedor =?1
order by closs_customer.nombre asc
`

func (q *Queries) GetCustomersBySalesman(ctx context.Context, code string) ([]ClossCustomer, error) {
	rows, err := q.db.QueryContext(ctx, getCustomersBySalesman, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossCustomer
	for rows.Next() {
		var i ClossCustomer
		if err := rows.Scan(
			&i.Codigo,
			&i.Nombre,
			&i.Email,
			&i.Direccion,
			&i.Telefonos,
			&i.Perscont,
			&i.Vendedor,
			&i.Contribespecial,
			&i.Status,
			&i.Sector,
			&i.Subsector,
			&i.Precio,
			&i.KneActiva,
			&i.KneMtomin,
			&i.Noemifac,
			&i.Noeminota,
			&i.Fchultvta,
			&i.Mtoultvta,
			&i.Prcdpagdia,
			&i.Promdiasp,
			&i.Riesgocrd,
			&i.Cantdocs,
			&i.Totmtodocs,
			&i.Prommtodoc,
			&i.Diasultvta,
			&i.Promdiasvta,
			&i.Limcred,
			&i.Fchcrea,
			&i.Dolarflete,
			&i.Nodolarflete,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE closs_customer SET 
    nombre = ?,
    direccion = ?,
    telefonos = ?,
    perscont = ?,
    vendedor = ?,
    contribespecial = ?,
    status = ?,
    sector = ?,
    subsector = ?,
    precio = ?,
    email = ?,
    kne_activa = ?,
    kne_mtomin = ?,
    noemifac = ?,
    noeminota = ?,
    fchultvta = ?,
    mtoultvta = ?,
    prcdpagdia = ?,
    promdiasp = ?,
    riesgocrd = ?,
    cantdocs = ?,
    totmtodocs = ?,
    prommtodoc = ?,
    diasultvta = ?,
    promdiasvta = ?,
    limcred = ?,
    dolarflete = ?,
    nodolarflete = ?,
    updated_at = ?
WHERE codigo = ?
RETURNING codigo, nombre, email, direccion, telefonos, perscont, vendedor, contribespecial, status, sector, subsector, precio, kne_activa, kne_mtomin, noemifac, noeminota, fchultvta, mtoultvta, prcdpagdia, promdiasp, riesgocrd, cantdocs, totmtodocs, prommtodoc, diasultvta, promdiasvta, limcred, fchcrea, dolarflete, nodolarflete, created_at, updated_at
`

type UpdateCustomerParams struct {
	Nombre          string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial int64
	Status          int64
	Sector          string
	Subsector       string
	Precio          int64
	Email           string
	KneActiva       int64
	KneMtomin       float64
	Noemifac        int64
	Noeminota       int64
	Fchultvta       string
	Mtoultvta       float64
	Prcdpagdia      float64
	Promdiasp       float64
	Riesgocrd       float64
	Cantdocs        int64
	Totmtodocs      float64
	Prommtodoc      float64
	Diasultvta      float64
	Promdiasvta     float64
	Limcred         float64
	Dolarflete      int64
	Nodolarflete    int64
	UpdatedAt       string
	Codigo          string
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (ClossCustomer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.Nombre,
		arg.Direccion,
		arg.Telefonos,
		arg.Perscont,
		arg.Vendedor,
		arg.Contribespecial,
		arg.Status,
		arg.Sector,
		arg.Subsector,
		arg.Precio,
		arg.Email,
		arg.KneActiva,
		arg.KneMtomin,
		arg.Noemifac,
		arg.Noeminota,
		arg.Fchultvta,
		arg.Mtoultvta,
		arg.Prcdpagdia,
		arg.Promdiasp,
		arg.Riesgocrd,
		arg.Cantdocs,
		arg.Totmtodocs,
		arg.Prommtodoc,
		arg.Diasultvta,
		arg.Promdiasvta,
		arg.Limcred,
		arg.Dolarflete,
		arg.Nodolarflete,
		arg.UpdatedAt,
		arg.Codigo,
	)
	var i ClossCustomer
	err := row.Scan(
		&i.Codigo,
		&i.Nombre,
		&i.Email,
		&i.Direccion,
		&i.Telefonos,
		&i.Perscont,
		&i.Vendedor,
		&i.Contribespecial,
		&i.Status,
		&i.Sector,
		&i.Subsector,
		&i.Precio,
		&i.KneActiva,
		&i.KneMtomin,
		&i.Noemifac,
		&i.Noeminota,
		&i.Fchultvta,
		&i.Mtoultvta,
		&i.Prcdpagdia,
		&i.Promdiasp,
		&i.Riesgocrd,
		&i.Cantdocs,
		&i.Totmtodocs,
		&i.Prommtodoc,
		&i.Diasultvta,
		&i.Promdiasvta,
		&i.Limcred,
		&i.Fchcrea,
		&i.Dolarflete,
		&i.Nodolarflete,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
