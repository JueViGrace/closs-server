package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Customer struct {
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          string    `json:"sector"`
	Subsector       string    `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       float64   `json:"kne_mtomin"`
	Noemifac        bool      `json:"noemifac"`
	Noeminota       bool      `json:"noeminota"`
	Fchultvta       time.Time `json:"fchultvta"`
	Mtoultvta       string    `json:"mtoultvta"`
	Prcdpagdia      string    `json:"prcdpagdia"`
	Promdiasp       string    `json:"promdiasp"`
	Riesgocrd       string    `json:"riesgocrd"`
	Cantdocs        int32     `json:"cantdocs"`
	Totmtodocs      string    `json:"totmtodocs"`
	Prommtodoc      string    `json:"prommtodoc"`
	Diasultvta      string    `json:"diasultvta"`
	Promdiasvta     string    `json:"promdiasvta"`
	Limcred         string    `json:"limcred"`
	Dolarflete      bool      `json:"dolarflete"`
	Nodolarflete    bool      `json:"nodolarflete"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type CreateCustomerRequest struct {
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          string    `json:"sector"`
	Subcodigo       string    `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       float64   `json:"kne_mtomin"`
	Noemifac        bool      `json:"noemifac"`
	Noeminota       bool      `json:"noeminota"`
	Fchultvta       time.Time `json:"fchultvta"`
	Mtoultvta       string    `json:"mtoultvta"`
	Prcdpagdia      string    `json:"prcdpagdia"`
	Promdiasp       string    `json:"promdiasp"`
	Riesgocrd       string    `json:"riesgocrd"`
	Cantdocs        int32     `json:"cantdocs"`
	Totmtodocs      string    `json:"totmtodocs"`
	Prommtodoc      string    `json:"prommtodoc"`
	Diasultvta      string    `json:"diasultvta"`
	Promdiasvta     string    `json:"promdiasvta"`
	Limcred         string    `json:"limcred"`
	Dolarflete      bool      `json:"dolarflete"`
	Nodolarflete    bool      `json:"nodolarflete"`
}

type UpdateCustomerRequest struct {
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          string    `json:"sector"`
	Subcodigo       string    `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       float64   `json:"kne_mtomin"`
	Noemifac        bool      `json:"noemifac"`
	Noeminota       bool      `json:"noeminota"`
	Fchultvta       time.Time `json:"fchultvta"`
	Mtoultvta       string    `json:"mtoultvta"`
	Prcdpagdia      string    `json:"prcdpagdia"`
	Promdiasp       string    `json:"promdiasp"`
	Riesgocrd       string    `json:"riesgocrd"`
	Cantdocs        int32     `json:"cantdocs"`
	Totmtodocs      string    `json:"totmtodocs"`
	Prommtodoc      string    `json:"prommtodoc"`
	Diasultvta      string    `json:"diasultvta"`
	Promdiasvta     string    `json:"promdiasvta"`
	Limcred         string    `json:"limcred"`
	Dolarflete      bool      `json:"dolarflete"`
	Nodolarflete    bool      `json:"nodolarflete"`
	ID              uuid.UUID `json:"id"`
}

func DbCustomerToCustomer(db *db.ClossCustomer) *Customer {
	return &Customer{
		Codigo:          db.Codigo,
		Nombre:          db.Nombre,
		Direccion:       db.Direccion,
		Telefonos:       db.Telefonos,
		Perscont:        db.Perscont,
		Vendedor:        db.Vendedor,
		Contribespecial: db.Contribespecial == 1,
		Status:          int16(db.Status),
		Sector:          db.Sector,
		Subsector:       db.Subsector,
		Precio:          int16(db.Precio),
		Email:           db.Email,
		KneActiva:       db.KneActiva == 1,
		KneMtomin:       db.KneMtomin,
		Noemifac:        db.Noemifac == 1,
		Noeminota:       db.Noeminota == 1,
		Fchultvta:       db.Fchultvta,
		Mtoultvta:       db.Mtoultvta,
		Prcdpagdia:      db.Prcdpagdia,
		Promdiasp:       db.Promdiasp,
		Riesgocrd:       db.Riesgocrd,
		Cantdocs:        db.Cantdocs,
		Totmtodocs:      db.Totmtodocs,
		Prommtodoc:      db.Prommtodoc,
		Diasultvta:      db.Diasultvta,
		Promdiasvta:     db.Promdiasvta,
		Limcred:         db.Limcred,
		Dolarflete:      db.Dolarflete,
		Nodolarflete:    db.Nodolarflete,
		CreatedAt:       db.CreatedAt,
		UpdatedAt:       db.UpdatedAt,
		DeletedAt:       db.DeletedAt.Time,
	}, nil
}

func CreateCustomerToDb(r *CreateCustomerRequest) (*db.CreateCustomerParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &db.CreateCustomerParams{
		ID:              id.String(),
		Codigo:          r.Codigo,
		Nombre:          r.Nombre,
		Direccion:       r.Direccion,
		Telefonos:       r.Telefonos,
		Perscont:        r.Perscont,
		Vendedor:        r.Vendedor,
		Contribespecial: r.Contribespecial,
		Status:          r.Status,
		Sector:          r.Sector,
		Subcodigo:       r.Subcodigo,
		Precio:          r.Precio,
		Email:           r.Email,
		KneActiva:       r.KneActiva,
		KneMtomin:       r.KneMtomin,
		Noemifac:        r.Noemifac,
		Noeminota:       r.Noeminota,
		Fchultvta:       r.Fchultvta,
		Mtoultvta:       r.Mtoultvta,
		Prcdpagdia:      r.Prcdpagdia,
		Promdiasp:       r.Promdiasp,
		Riesgocrd:       r.Riesgocrd,
		Cantdocs:        r.Cantdocs,
		Totmtodocs:      r.Totmtodocs,
		Prommtodoc:      r.Prommtodoc,
		Diasultvta:      r.Diasultvta,
		Promdiasvta:     r.Promdiasvta,
		Limcred:         r.Limcred,
		Dolarflete:      r.Dolarflete,
		Nodolarflete:    r.Nodolarflete,
	}, nil
}

func UpdateCustomerToDb(r *UpdateCustomerRequest) *db.UpdateCustomerParams {
	return &db.UpdateCustomerParams{
		Codigo:          r.Codigo,
		Nombre:          r.Nombre,
		Direccion:       r.Direccion,
		Telefonos:       r.Telefonos,
		Perscont:        r.Perscont,
		Vendedor:        r.Vendedor,
		Contribespecial: r.Contribespecial,
		Status:          r.Status,
		Sector:          r.Sector,
		Subcodigo:       r.Subcodigo,
		Precio:          r.Precio,
		Email:           r.Email,
		KneActiva:       r.KneActiva,
		KneMtomin:       r.KneMtomin,
		Noemifac:        r.Noemifac,
		Noeminota:       r.Noeminota,
		Fchultvta:       r.Fchultvta,
		Mtoultvta:       r.Mtoultvta,
		Prcdpagdia:      r.Prcdpagdia,
		Promdiasp:       r.Promdiasp,
		Riesgocrd:       r.Riesgocrd,
		Cantdocs:        r.Cantdocs,
		Totmtodocs:      r.Totmtodocs,
		Prommtodoc:      r.Prommtodoc,
		Diasultvta:      r.Diasultvta,
		Promdiasvta:     r.Promdiasvta,
		Limcred:         r.Limcred,
		Dolarflete:      r.Dolarflete,
		Nodolarflete:    r.Nodolarflete,
		ID:              r.ID.String(),
	}
}
