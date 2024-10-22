package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type CompanyStore interface {
	GetCompanies()
	GetCompany(id uuid.UUID)
	CreateCompany()
	UpdateCompany(id uuid.UUID)
	DeleteCompany(id uuid.UUID)
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

func (s *companyStore) GetCompanies() {
}

func (s *companyStore) GetCompany(id uuid.UUID) {
}

func (s *companyStore) CreateCompany() {
}

func (s *companyStore) UpdateCompany(id uuid.UUID) {
}

func (s *companyStore) DeleteCompany(id uuid.UUID) {
}
