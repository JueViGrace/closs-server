package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/gofiber/fiber/v2"
)

type SalesmanHandler interface {
	GetSalesmen(c *fiber.Ctx) error
	GetSalesman(c *fiber.Ctx) error
	CreateSalesman(c *fiber.Ctx) error
	UpdateSalesman(c *fiber.Ctx) error
	DeleteSalesman(c *fiber.Ctx) error
}

type salesmanHandler struct {
	db data.SalesmanStore
}

func NewSalesmanHandler(db data.SalesmanStore) SalesmanHandler {
	return &salesmanHandler{
		db: db,
	}
}

func (h *salesmanHandler) GetSalesmen(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	salesmen, err := h.db.GetSalesmen()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(salesmen, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *salesmanHandler) GetSalesman(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	salesman, err := h.db.GetSalesman(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(salesman, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *salesmanHandler) CreateSalesman(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateSalesmanRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateSalesman(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *salesmanHandler) UpdateSalesman(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateSalesmanRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateSalesman(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *salesmanHandler) DeleteSalesman(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.DeleteSalesman(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}
