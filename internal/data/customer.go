package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type CustomerStore interface {
	GetCustomers() ([]*types.Customer, error)
	GetCustomer(id uuid.UUID) (*types.Customer, error)
	CreateCustomer(r *types.CreateCustomerRequest) (string, error)
	UpdateCustomer(r *types.UpdateCustomerRequest) (string, error)
	DeleteCustomer(id uuid.UUID) error
}

func (s *storage) CustomerStore() CustomerStore {
	return NewCustomerStore(s.ctx, s.queries)
}

type customerStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewCustomerStore(ctx context.Context, db *db.Queries) CustomerStore {
	return &customerStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *customerStore) GetCustomers() ([]*types.Customer, error) {
	return nil, nil
}

func (s *customerStore) GetCustomer(id uuid.UUID) (*types.Customer, error) {
	return nil, nil
}

func (s *customerStore) CreateCustomer(r *types.CreateCustomerRequest) (string, error) {
	return "", nil
}

func (s *customerStore) UpdateCustomer(r *types.UpdateCustomerRequest) (string, error) {
	return "", nil
}

func (s *customerStore) DeleteCustomer(id uuid.UUID) error {
	return nil
}
