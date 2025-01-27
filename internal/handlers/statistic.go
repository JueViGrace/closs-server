package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type StatisticHandler interface {
	GetStatistics(c *fiber.Ctx, d *types.AuthData) error
	GetStatisticBySalesman(c *fiber.Ctx, d *types.AuthData) error
	CreateStatistic(c *fiber.Ctx) error
	UpdateStatistic(c *fiber.Ctx) error
}

type statisticHandler struct {
	db data.StatisticStore
}

func NewStatisticHandler(db data.StatisticStore) StatisticHandler {
	return &statisticHandler{
		db: db,
	}
}

// todo: use auth data for search
func (h *statisticHandler) GetStatistics(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)

	statistics, err := h.db.GetStatistics()
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(statistics, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *statisticHandler) GetStatisticBySalesman(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)
	statistic, err := h.db.GetStatisticBySalesman(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(statistic, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *statisticHandler) CreateStatistic(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateStatisticRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateStatistic(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *statisticHandler) UpdateStatistic(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateStatisticRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateStatistic(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}
