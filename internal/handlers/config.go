package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type ConfigHandler interface {
	GetConfigs(c *fiber.Ctx) error
	GetConfig(c *fiber.Ctx) error
	CreateConfig(c *fiber.Ctx) error
	UpdateConfig(c *fiber.Ctx) error
	DeleteConfig(c *fiber.Ctx) error
}

type configHandler struct {
	db data.ConfigStore
}

func NewConfigHandler(db data.ConfigStore) ConfigHandler {
	return &configHandler{
		db: db,
	}
}

func (h *configHandler) GetConfigs(c *fiber.Ctx) error {
	return nil
}

func (h *configHandler) GetConfig(c *fiber.Ctx) error {
	return nil
}

func (h *configHandler) CreateConfig(c *fiber.Ctx) error {
	return nil
}

func (h *configHandler) UpdateConfig(c *fiber.Ctx) error {
	return nil
}

func (h *configHandler) DeleteConfig(c *fiber.Ctx) error {
	return nil
}
