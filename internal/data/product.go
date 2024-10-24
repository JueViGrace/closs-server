package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type ProductStore interface {
	GetProducts() (products []*types.Product, err error)
	GetProduct(id *uuid.UUID) (product *types.Product, err error)
	CreateProduct(r *types.CreateProductRequest) (msg string, err error)
	UpdateProduct(r *types.UpdateProductRequest) (msg string, err error)
	DeleteProduct(id *uuid.UUID) (err error)
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
	products := make([]*types.Product, 0)

	dbProducts, err := s.db.GetAllProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbProduct := range dbProducts {
		product, err := types.DbProductToProduct(&dbProduct)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *productStore) GetProduct(id *uuid.UUID) (*types.Product, error) {
	product := new(types.Product)

	dbProduct, err := s.db.GetProductById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	product, err = types.DbProductToProduct(&dbProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productStore) CreateProduct(r *types.CreateProductRequest) (string, error) {
	cr, err := types.CreateProductToDb(r)
	if err != nil {
		return "", err
	}

	err = s.db.CreateProduct(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *productStore) UpdateProduct(r *types.UpdateProductRequest) (string, error) {
	ur := types.UpdateProductToDb(r)

	err := s.db.UpdateProduct(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updates!", nil
}

func (s *productStore) DeleteProduct(id *uuid.UUID) error {
	err := s.db.SoftDeleteProduct(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}
