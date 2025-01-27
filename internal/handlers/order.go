package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx, d *types.AuthData) error
	GetOrderByCode(c *fiber.Ctx, d *types.AuthData) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
}

type orderHandler struct {
	db data.OrderStore
}

func NewOrderHandler(db data.OrderStore) OrderHandler {
	return &orderHandler{
		db: db,
	}
}

// todo: use auth data for search
func (h *orderHandler) GetOrders(c *fiber.Ctx, d *types.AuthData) error {
	return nil
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx, d *types.AuthData) error {
	return nil
}

func (h *orderHandler) GetAllOrdersByManager(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrdersByManager(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetAllOrdersBySalesman(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrdersBySalesman(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetAllOrdersByCustomer(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrdersByCustomer(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) UpdateOrder(c *fiber.Ctx) error {
	return nil
}
