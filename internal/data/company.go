package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type CompanyStore interface {
	AdminGetCompanies() (companies []*types.Company, err error)
	AdminGetCompanyById(id *uuid.UUID) (company *types.Company, err error)
	CreateCompany(r *types.CreateCompanyRequest) (msg string, err error)
	UpdateCompany(r *types.UpdateCompanyRequest) (msg string, err error)
	DeleteCompany(id *uuid.UUID) (err error)
	GetCompanyByCode(code string) (company *types.Company, err error)
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

func (s *companyStore) AdminGetCompanies() ([]*types.Company, error) {
	companies := make([]*types.Company, 0)

	dbCompanies, err := s.db.AdminGetCompanies(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbCompany := range dbCompanies {
		company, err := types.DbComanyToCompany(&dbCompany)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (s *companyStore) AdminGetCompanyById(id *uuid.UUID) (*types.Company, error) {
	company := new(types.Company)

	dbCompany, err := s.db.AdminGetCompanyById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	company, err = types.DbComanyToCompany(&dbCompany)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyStore) CreateCompany(r *types.CreateCompanyRequest) (string, error) {
	cr, err := types.CreateCompanyToDb(r)
	if err != nil {
		return "", err
	}

	err = s.db.CreateCompany(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *companyStore) UpdateCompany(r *types.UpdateCompanyRequest) (string, error) {
	ur := types.UpdateCompanyToDb(r)

	err := s.db.UpdateCompany(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updated!", nil
}

func (s *companyStore) DeleteCompany(id *uuid.UUID) error {
	err := s.db.SoftDeleteCompany(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}

func (s *companyStore) GetCompanyByCode(code string) (*types.Company, error) {
	company := new(types.Company)

	dbCompany, err := s.db.GetCompanyByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	company, err = types.DbComanyToCompany(&dbCompany)
	if err != nil {
		return nil, err
	}

	return company, nil
}
