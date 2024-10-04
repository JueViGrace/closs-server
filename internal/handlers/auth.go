package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}

type authHandler struct {
	db data.AuthStore
}

func NewAuthHandler(db data.AuthStore) AuthHandler {
	return &authHandler{
		db: db,
	}
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	res := types.RespondOk("Sign in handler", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) SignUp(c *fiber.Ctx) error {
	res := types.RespondOk("Sign up handler", "Success")
	return c.Status(res.Status).JSON(res)
}
