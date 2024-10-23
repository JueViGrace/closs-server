package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Customer struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          int32     `json:"sector"`
	Subcodigo       int32     `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       string    `json:"kne_mtomin"`
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
	Sector          int32     `json:"sector"`
	Subcodigo       int32     `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       string    `json:"kne_mtomin"`
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
	Sector          int32     `json:"sector"`
	Subcodigo       int32     `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       string    `json:"kne_mtomin"`
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

func DbCustomerToCustomer(cus *db.Cliente) (*Customer, error) {
	userId := new(uuid.UUID)

	id, err := uuid.Parse(cus.ID)
	if err != nil {
		return nil, err
	}

	if cus.UserID.Valid {
		*userId, err = uuid.Parse(cus.UserID.String)
		if err != nil {
			return nil, err
		}
	}

	return &Customer{
		ID:              id,
		UserID:          *userId,
		Codigo:          cus.Codigo,
		Nombre:          cus.Nombre,
		Direccion:       cus.Direccion,
		Telefonos:       cus.Telefonos,
		Perscont:        cus.Perscont,
		Vendedor:        cus.Vendedor,
		Contribespecial: cus.Contribespecial,
		Status:          cus.Status,
		Sector:          cus.Sector,
		Subcodigo:       cus.Subcodigo,
		Precio:          cus.Precio,
		Email:           cus.Email,
		KneActiva:       cus.KneActiva,
		KneMtomin:       cus.KneMtomin,
		Noemifac:        cus.Noemifac,
		Noeminota:       cus.Noeminota,
		Fchultvta:       cus.Fchultvta,
		Mtoultvta:       cus.Mtoultvta,
		Prcdpagdia:      cus.Prcdpagdia,
		Promdiasp:       cus.Promdiasp,
		Riesgocrd:       cus.Riesgocrd,
		Cantdocs:        cus.Cantdocs,
		Totmtodocs:      cus.Totmtodocs,
		Prommtodoc:      cus.Prommtodoc,
		Diasultvta:      cus.Diasultvta,
		Promdiasvta:     cus.Promdiasvta,
		Limcred:         cus.Limcred,
		Dolarflete:      cus.Dolarflete,
		Nodolarflete:    cus.Nodolarflete,
		CreatedAt:       cus.CreatedAt,
		UpdatedAt:       cus.UpdatedAt,
		DeletedAt:       cus.DeletedAt.Time,
	}, nil
}

func CreateCustomerToDb(r *CreateCustomerRequest) (*db.Cliente, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &db.Cliente{
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
