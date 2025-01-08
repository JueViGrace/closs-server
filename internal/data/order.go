package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type OrderStore interface {
	GetOrders() (orders []types.OrderResponse, err error)
	GetOrdersWithLines() (orders []types.OrderWithLinesResponse, err error)
	GetOrderByCode(code string) (order *types.OrderResponse, err error)
	GetOrderByCodeWithLines(code string) (order *types.OrderWithLinesResponse, err error)
	GetAllOrdersByManager(code string) (orders []types.OrderResponse, err error)
	GetOrdersByManager(code string) (orders []types.OrderResponse, err error)
	GetAllOrdersBySalesman(code string) (orders []types.OrderResponse, err error)
	GetOrdersBySalesman(code string) (orders []types.OrderResponse, err error)
	GetAllOrdersByCustomer(code string) (orders []types.OrderResponse, err error)
	GetOrdersByCustomer(code string) (orders []types.OrderResponse, err error)
	CreateOrder(r *types.CreateOrderRequest) (order *types.OrderWithLinesResponse, err error)
	UpdateOrder(r *types.UpdateOrderRequest) (order *types.OrderWithLinesResponse, err error)
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

func (s *orderStore) GetOrders() ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrdersWithLines() ([]types.OrderWithLinesResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrderByCode(code string) (*types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrderByCodeWithLines(code string) (*types.OrderWithLinesResponse, error) {
	return nil, nil
}

func (s *orderStore) GetAllOrdersByManager(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrdersByManager(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetAllOrdersBySalesman(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrdersBySalesman(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetAllOrdersByCustomer(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) GetOrdersByCustomer(code string) ([]types.OrderResponse, error) {
	return nil, nil
}

func (s *orderStore) CreateOrder(r *types.CreateOrderRequest) (*types.OrderWithLinesResponse, error) {
	return nil, nil
}

func (s *orderStore) UpdateOrder(r *types.UpdateOrderRequest) (*types.OrderWithLinesResponse, error) {
	return nil, nil
}

