package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
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
	return nil
}

func (h *salesmanHandler) GetSalesman(c *fiber.Ctx) error {
	return nil
}

func (h *salesmanHandler) CreateSalesman(c *fiber.Ctx) error {
	return nil
}

func (h *salesmanHandler) UpdateSalesman(c *fiber.Ctx) error {
	return nil
}

func (h *salesmanHandler) DeleteSalesman(c *fiber.Ctx) error {
	return nil
}
