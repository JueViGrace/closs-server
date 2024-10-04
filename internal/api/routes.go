package api

import (
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (a *api) RegisterRoutes() {
	a.HealthRoute()
	a.MonitorRoute()

	a.AuthRoutes()
}

func (a *api) HealthRoute() {
	a.Get("/api/health", func(c *fiber.Ctx) error {
		res := types.RespondOk(a.db.Health(), "Success")
		return c.Status(res.Status).JSON(res)
	})
}

func (a *api) MonitorRoute() {
	a.Get("/api/metrics", monitor.New())
}
