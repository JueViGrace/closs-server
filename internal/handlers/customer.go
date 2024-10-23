package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	GetCustomers(c *fiber.Ctx) error
	GetCustomer(c *fiber.Ctx) error
	CreateCustomer(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}

type customerHandler struct {
	db data.CustomerStore
}

func NewCustomerHandler(db data.CustomerStore) CustomerHandler {
	return &customerHandler{
		db: db,
	}
}

func (h *customerHandler) GetCustomers(c *fiber.Ctx) error {
	return nil
}

func (h *customerHandler) GetCustomer(c *fiber.Ctx) error {
	return nil
}

func (h *customerHandler) CreateCustomer(c *fiber.Ctx) error {
	return nil
}

func (h *customerHandler) UpdateCustomer(c *fiber.Ctx) error {
	return nil
}

func (h *customerHandler) DeleteCustomer(c *fiber.Ctx) error {
	return nil
}
