package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx, d *types.AuthData) error
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
	res := new(types.APIResponse)

	users, err := h.db.GetUsers()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(users, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetUserById(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)

	user, err := h.db.GetUserById(d.UserId)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}
