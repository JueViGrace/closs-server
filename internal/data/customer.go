package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type CustomerStore interface {
	GetCustomers()
	GetCustomer(id uuid.UUID)
	CreateCustomer()
	UpdateCustomer(id uuid.UUID)
	DeleteCustomer(id uuid.UUID)
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

func (s *customerStore) GetCustomers() {
}

func (s *customerStore) GetCustomer(id uuid.UUID) {
}

func (s *customerStore) CreateCustomer() {
}

func (s *customerStore) UpdateCustomer(id uuid.UUID) {
}

func (s *customerStore) DeleteCustomer(id uuid.UUID) {
}

