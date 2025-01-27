package handlers

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx, d *types.AuthData) error
	RecoverPassword(c *fiber.Ctx) error
}

type authHandler struct {
	db data.AuthStore
}

func NewAuthHandler(db data.AuthStore) AuthHandler {
	return &authHandler{
		db: db,
	}
}

// todo: add validation
func (h *authHandler) SignIn(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.SignInRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignIn(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	c.Cookie(&fiber.Cookie{
		Name:    "closs_access_token",
		Value:   token.AccessToken,
		Expires: time.Now().Add(1 * time.Hour),
	})
	c.Cookie(&fiber.Cookie{
		Name:    "closs_refresh_token",
		Value:   token.RefreshToken,
		Expires: time.Now().Add(24 * time.Hour),
	})

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) Refresh(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)
	r := new(types.RefreshRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.Refresh(r, d)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	c.ClearCookie("closs_access_token")
	c.ClearCookie("closs_refresh_token")
	c.Cookie(&fiber.Cookie{
		Name:    "closs_access_token",
		Value:   token.AccessToken,
		Expires: time.Now().Add(1 * time.Hour),
	})
	c.Cookie(&fiber.Cookie{
		Name:    "closs_refresh_token",
		Value:   token.RefreshToken,
		Expires: time.Now().Add(24 * time.Hour),
	})

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	res := types.RespondOk("RecoverPassword handler", "Success")
	return c.Status(res.Status).JSON(res)
}
