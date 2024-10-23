package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userHandler struct {
	db data.UserStore
}

func NewUserHandler(db data.UserStore) UserHandler {
	return &userHandler{
		db: db,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	return nil
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	return nil
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	return nil
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	return nil
}
