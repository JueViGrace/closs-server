package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type CustomerStore interface {
	AdminGetCustomers() (customers []*types.Customer, err error)
	AdminGetCustomer(id *uuid.UUID) (customer *types.Customer, err error)
	CreateCustomer(r *types.CreateCustomerRequest) (msg string, err error)
	UpdateCustomer(r *types.UpdateCustomerRequest) (msg string, err error)
	DeleteCustomer(id *uuid.UUID) (err error)
	GetCustomersByManager(code string) (customers []*types.Customer, err error)
	GetCustomersBySalesman(code string) (customers []*types.Customer, err error)
	GetCustomerById(id *uuid.UUID) (customer *types.Customer, err error)
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

func (s *customerStore) AdminGetCustomers() ([]*types.Customer, error) {
	customers := make([]*types.Customer, 0)

	dbCustomers, err := s.db.AdminGetCustomers(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbCustomer := range dbCustomers {
		customer, err := types.DbCustomerToCustomer(&dbCustomer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (s *customerStore) AdminGetCustomer(id *uuid.UUID) (*types.Customer, error) {
	customer := new(types.Customer)

	dbCustomer, err := s.db.GetCustomerById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	customer, err = types.DbCustomerToCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerStore) CreateCustomer(r *types.CreateCustomerRequest) (string, error) {
	cr, err := types.CreateCustomerToDb(r)
	if err != nil {
		return "", err
	}

	err = s.db.CreateCustomer(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *customerStore) UpdateCustomer(r *types.UpdateCustomerRequest) (string, error) {
	ur := types.UpdateCustomerToDb(r)

	err := s.db.UpdateCustomer(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updated!", nil
}

func (s *customerStore) DeleteCustomer(id *uuid.UUID) error {
	err := s.db.SoftDeleteCustomer(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}

func GetCustomersByManager(code string) ([]*types.Customer, error) {
}

func GetCustomersBySalesman(code string) ([]*types.Customer, error) {
}

func GetCustomerById(id *uuid.UUID) (*types.Customer, error) {
}
