package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/gofiber/fiber/v2"
)

type CompanyHandler interface {
	GetCompanies(c *fiber.Ctx) error
	GetCompany(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
	GetCompanyByCode(c *fiber.Ctx) error
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
	res := new(types.APIResponse)

	companies, err := h.db.AdminGetCompanies()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(companies, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) GetCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	company, err := h.db.AdminGetCompanyById(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(company, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) CreateCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateCompanyRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateCompany(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) UpdateCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateCompanyRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateCompany(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) DeleteCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.DeleteCompany(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) GetCompanyByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	code := c.Params("code")

	company, err := h.db.GetCompanyByCode(code)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(company, "Success")
	return c.Status(res.Status).JSON(res)
}
