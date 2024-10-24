package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type SalesmanStore interface {
	GetSalesmen() (salesmen []*types.Salesman, err error)
	GetSalesman(id *uuid.UUID) (salesman *types.Salesman, err error)
	CreateSalesman(r *types.CreateSalesmanRequest) (msg string, err error)
	UpdateSalesman(r *types.UpdateSalesmanRequest) (msg string, err error)
	DeleteSalesman(id *uuid.UUID) (err error)
}

func (s *storage) SalesmanStore() SalesmanStore {
	return NewSalesmanStore(s.ctx, s.queries)
}

type salesmanStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewSalesmanStore(ctx context.Context, db *db.Queries) SalesmanStore {
	return &salesmanStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *salesmanStore) GetSalesmen() ([]*types.Salesman, error) {
	salesmen := make([]*types.Salesman, 0)

	dbSalesmen, err := s.db.AdminGetSalesmen(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbSalesman := range dbSalesmen {
		salesman, err := types.DbSalesmanToSalesman(&dbSalesman)
		if err != nil {
			return nil, err
		}

		salesmen = append(salesmen, salesman)
	}

	return salesmen, nil
}

func (s *salesmanStore) GetSalesman(id *uuid.UUID) (*types.Salesman, error) {
	salesman := new(types.Salesman)

	dbSalesman, err := s.db.GetSalesmanById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	salesman, err = types.DbSalesmanToSalesman(&dbSalesman)
	if err != nil {
		return nil, err
	}

	return salesman, nil
}

func (s *salesmanStore) CreateSalesman(r *types.CreateSalesmanRequest) (string, error) {
	cr, err := types.CreateSalesmanToDb(r)
	if err != nil {
		return "", nil
	}

	err = s.db.CreateSalesman(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *salesmanStore) UpdateSalesman(r *types.UpdateSalesmanRequest) (string, error) {
	ur := types.UpdateSalesmanToDb(r)

	err := s.db.UpdateSalesman(s.ctx, *ur)
	if err != nil {
		return "", nil
	}

	return "Updated!", nil
}

func (s *salesmanStore) DeleteSalesman(id *uuid.UUID) error {
	err := s.db.SoftDeleteSalesman(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}
