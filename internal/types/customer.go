package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Customer struct {
	ID              string
	UserID          string
	Codigo          string
	Nombre          string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial bool
	Status          int16
	Sector          int32
	Subcodigo       int32
	Precio          int16
	Email           string
	KneActiva       bool
	KneMtomin       string
	Noemifac        bool
	Noeminota       bool
	Fchultvta       time.Time
	Mtoultvta       string
	Prcdpagdia      string
	Promdiasp       string
	Riesgocrd       string
	Cantdocs        int32
	Totmtodocs      string
	Prommtodoc      string
	Diasultvta      string
	Promdiasvta     string
	Limcred         string
	Dolarflete      bool
	Nodolarflete    bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type CreateCustomerRequest struct {
	Codigo          string
	Nombre          string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial bool
	Status          int16
	Sector          int32
	Subcodigo       int32
	Precio          int16
	Email           string
	KneActiva       bool
	KneMtomin       string
	Noemifac        bool
	Noeminota       bool
	Fchultvta       time.Time
	Mtoultvta       string
	Prcdpagdia      string
	Promdiasp       string
	Riesgocrd       string
	Cantdocs        int32
	Totmtodocs      string
	Prommtodoc      string
	Diasultvta      string
	Promdiasvta     string
	Limcred         string
	Dolarflete      bool
	Nodolarflete    bool
}

type UpdateCustomerRequest struct {
	Codigo          string
	Nombre          string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial bool
	Status          int16
	Sector          int32
	Subcodigo       int32
	Precio          int16
	Email           string
	KneActiva       bool
	KneMtomin       string
	Noemifac        bool
	Noeminota       bool
	Fchultvta       time.Time
	Mtoultvta       string
	Prcdpagdia      string
	Promdiasp       string
	Riesgocrd       string
	Cantdocs        int32
	Totmtodocs      string
	Prommtodoc      string
	Diasultvta      string
	Promdiasvta     string
	Limcred         string
	Dolarflete      bool
	Nodolarflete    bool
	ID              string
}

func DbCustomerToCustomer(cus *db.Cliente) (*Customer, error) {
	id, err := uuid.Parse(cus.ID)
	if err != nil {
		return nil, err
	}

	return &Customer{
		ID: id,
		UserID: cus.UserID,
		Codigo: cus.Codigo,
		Nombre: cus.Nombre,
		Direccion: cus.Direccion,
		Telefonos: cus.Telefonos,
		Perscont: cus.Perscont,
		Vendedor: cus.Vendedor,
		Contribespecial: cus.Contribespecial,
		Status: cus.Status,
		Sector: cus.Sector,
		Subcodigo: cus.Subcodigo,
		Precio: cus.Precio,
		Email: cus.Email,
		KneActiva: cus.KneActiva,
		KneMtomin: cus.KneMtomin,
		Noemifac: cus.Noemifac,
		Noeminota: cus.Noeminota,
		Fchultvta: cus.Fchultvta,
		Mtoultvta: cus.Mtoultvta,
		Prcdpagdia: cus.Prcdpagdia,
		Promdiasp: cus.Promdiasp,
		Riesgocrd: cus.Riesgocrd,
		Cantdocs: cus.Cantdocs,
		Totmtodocs: cus.Totmtodocs,
		Prommtodoc: cus.Prommtodoc,
		Diasultvta: cus.Diasultvta,
		Promdiasvta: cus.Promdiasvta,
		Limcred: cus.Limcred,
		Dolarflete: cus.Dolarflete,
		Nodolarflete: cus.Nodolarflete,
		CreatedAt: cus.CreatedAt,
		UpdatedAt: cus.UpdatedAt,
		DeletedAt: cus.DeletedAt,
	}, nil
}
