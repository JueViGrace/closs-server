package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type productHandler struct {
	db data.ProductStore
}

func NewProductHandler(db data.ProductStore) ProductHandler {
	return &productHandler{
		db: db,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	return nil
}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	return nil
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	return nil
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	return nil
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	return nil
}
