package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type ProductStore interface {
	GetProducts() ([]*types.Product, error)
	GetProduct(id uuid.UUID) (*types.Product, error)
	CreateProduct(r *types.CreateProductRequest) (string, error)
	UpdateProduct(r *types.UpdateProductRequest) (string, error)
	DeleteProduct(id uuid.UUID) error
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

func (s *productStore) GetProducts() ([]*types.Product, error) {
	return nil, nil
}

func (s *productStore) GetProduct(id uuid.UUID) (*types.Product, error) {
	return nil, nil
}

func (s *productStore) CreateProduct(r *types.CreateProductRequest) (string, error) {
	return "", nil
}

func (s *productStore) UpdateProduct(r *types.UpdateProductRequest) (string, error) {
	return "", nil
}

func (s *productStore) DeleteProduct(id uuid.UUID) error {
	return nil
}
