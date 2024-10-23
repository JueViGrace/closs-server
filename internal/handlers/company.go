package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type CompanyHandler interface {
	GetCompanies(c *fiber.Ctx) error
	GetCompany(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
}

type companyHandler struct {
	db data.CompanyStore
}

func NewCompanyHandler(db data.CompanyStore) CompanyHandler {
	return &companyHandler{
		db: db,
	}
}

func (h *companyHandler) GetCompanies(c *fiber.Ctx) error {
	return nil
}

func (h *companyHandler) GetCompany(c *fiber.Ctx) error {
	return nil
}

func (h *companyHandler) CreateCompany(c *fiber.Ctx) error {
	return nil
}

func (h *companyHandler) UpdateCompany(c *fiber.Ctx) error {
	return nil
}

func (h *companyHandler) DeleteCompany(c *fiber.Ctx) error {
	return nil
}
