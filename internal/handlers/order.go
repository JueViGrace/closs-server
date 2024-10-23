package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
}

type orderHandler struct {
	db data.OrderStore
}

func NewOrderHandler(db data.OrderStore) OrderHandler {
	return &orderHandler{
		db: db,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrder(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) UpdateOrder(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) DeleteOrder(c *fiber.Ctx) error {
	return nil
}
