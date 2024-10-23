package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type StatisticHandler interface {
	GetStatistics(c *fiber.Ctx) error
	GetStatistic(c *fiber.Ctx) error
	CreateStatistic(c *fiber.Ctx) error
	UpdateStatistic(c *fiber.Ctx) error
	DeleteStatistic(c *fiber.Ctx) error
}

type statisticHandler struct {
	db data.StatisticStore
}

func NewStatisticHandler(db data.StatisticStore) StatisticHandler {
	return &statisticHandler{
		db: db,
	}
}

func (h *statisticHandler) GetStatistics(c *fiber.Ctx) error {
	return nil
}

func (h *statisticHandler) GetStatistic(c *fiber.Ctx) error {
	return nil
}

func (h *statisticHandler) CreateStatistic(c *fiber.Ctx) error {
	return nil
}

func (h *statisticHandler) UpdateStatistic(c *fiber.Ctx) error {
	return nil
}

func (h *statisticHandler) DeleteStatistic(c *fiber.Ctx) error {
	return nil
}
