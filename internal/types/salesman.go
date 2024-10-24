package types

import "github.com/JueViGrace/clo-backend/internal/db"

type Salesman struct{}

type CreateSalesmanRequest struct{}

type UpdateSalesmanRequest struct{}

func DbSalesmanToSalesman(db *db.Vendedor) (*Salesman, error) {
	return &Salesman{}, nil
}

func CreateSalesmanToDb(r *CreateSalesmanRequest) (*db.CreateSalesmanParams, error) {
	return &db.CreateSalesmanParams{}, nil
}

func UpdateSalesmanToDb(r *UpdateSalesmanRequest) *db.UpdateSalesmanParams {
	return &db.UpdateSalesmanParams{}
}
