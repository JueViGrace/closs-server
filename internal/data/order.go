package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type OrderStore interface {
	GetOrders()
	GetOrder(id uuid.UUID)
	CreateOrder()
	UpdateOrder(id uuid.UUID)
	DeleteOrder(id uuid.UUID)
}

func (s *storage) OrderStore() OrderStore {
	return NewOrderStore(s.ctx, s.queries)
}

type orderStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewOrderStore(ctx context.Context, db *db.Queries) OrderStore {
	return &orderStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *orderStore) GetOrders() {
}

func (s *orderStore) GetOrder(id uuid.UUID) {
}

func (s *orderStore) CreateOrder() {
}

func (s *orderStore) UpdateOrder(id uuid.UUID) {
}

func (s *orderStore) DeleteOrder(id uuid.UUID) {
}