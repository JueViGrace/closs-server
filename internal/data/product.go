package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type ProductStore interface {
	GetProducts()
	GetProduct(id uuid.UUID)
	CreateProduct()
	UpdateProduct(id uuid.UUID)
	DeleteProduct(id uuid.UUID)
}

func (s *storage) ProductStore() ProductStore {
	return NewProductStore(s.ctx, s.queries)
}

type productStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewProductStore(ctx context.Context, db *db.Queries) ProductStore {
	return &productStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *productStore) GetProducts() {
}

func (s *productStore) GetProduct(id uuid.UUID) {
}

func (s *productStore) CreateProduct() {
}

func (s *productStore) UpdateProduct(id uuid.UUID) {
}

func (s *productStore) DeleteProduct(id uuid.UUID) {
}
