package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type SalesmanStore interface {
	GetSalesmen()
	GetSalesman(id uuid.UUID)
	CreateSalesman()
	UpdateSalesman(id uuid.UUID)
	DeleteSalesman(id uuid.UUID)
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

func (s *salesmanStore) GetSalesmen() {
}

func (s *salesmanStore) GetSalesman(id uuid.UUID) {
}

func (s *salesmanStore) CreateSalesman() {
}

func (s *salesmanStore) UpdateSalesman(id uuid.UUID) {
}

func (s *salesmanStore) DeleteSalesman(id uuid.UUID) {
}
