package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type CompanyStore interface {
	GetCompanies() ([]*types.Company, error)
	GetCompany(id uuid.UUID) (*types.Company, error)
	CreateCompany(r *types.CreateCompanyRequest) (string, error)
	UpdateCompany(r *types.UpdateCompanyRequest) (string, error)
	DeleteCompany(id uuid.UUID) error
}

func (s *storage) CompanyStore() CompanyStore {
	return NewCompanyStore(s.ctx, s.queries)
}

type companyStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewCompanyStore(ctx context.Context, db *db.Queries) CompanyStore {
	return &companyStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *companyStore) GetCompanies() ([]*types.Company, error) {
	return nil, nil
}

func (s *companyStore) GetCompany(id uuid.UUID) (*types.Company, error) {
	return nil, nil
}

func (s *companyStore) CreateCompany(r *types.CreateCompanyRequest) (string, error) {
	return "", nil
}

func (s *companyStore) UpdateCompany(r *types.UpdateCompanyRequest) (string, error) {
	return "", nil
}

func (s *companyStore) DeleteCompany(id uuid.UUID) error {
	return nil
}
