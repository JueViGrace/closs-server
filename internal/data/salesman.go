package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type SalesmanStore interface {
	GetSalesmen() ([]*types.Salesman, error)
	GetSalesman(id uuid.UUID) (*types.Salesman, error)
	CreateSalesman(r *types.CreateSalesmanRequest) (string, error)
	UpdateSalesman(r *types.UpdateSalesmanRequest) (string, error)
	DeleteSalesman(id uuid.UUID) error
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
	return nil, nil
}

func (s *salesmanStore) GetSalesman(id uuid.UUID) (*types.Salesman, error) {
	return nil, nil
}

func (s *salesmanStore) CreateSalesman(r *types.CreateSalesmanRequest) (string, error) {
	return "", nil
}

func (s *salesmanStore) UpdateSalesman(r *types.UpdateSalesmanRequest) (string, error) {
	return "", nil
}

func (s *salesmanStore) DeleteSalesman(id uuid.UUID) error {
	return nil
}
