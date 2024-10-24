package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/gofiber/fiber/v2"
)

type ConfigHandler interface {
	GetConfigs(c *fiber.Ctx) error
	GetConfig(c *fiber.Ctx) error
	CreateConfig(c *fiber.Ctx) error
	UpdateConfig(c *fiber.Ctx) error
	DeleteConfig(c *fiber.Ctx) error
	GetConfigByUser(c *fiber.Ctx) error
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
	res := new(types.APIResponse)

	configs, err := h.db.AdminGetCofigs()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(configs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) GetConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	config, err := h.db.AdminGetConfigById(id)
	if err != nil {

		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(config, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) CreateConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateConfigRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateConfig(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) UpdateConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateConfigRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateConfig(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) DeleteConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.DeleteConfig(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) GetConfigByUser(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id, err := util.GetIdFromParams(c.Params(":id"))
	if err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	configs, err := h.db.GetConfigByUser(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(configs, "Success")
	return c.Status(res.Status).JSON(res)
}
