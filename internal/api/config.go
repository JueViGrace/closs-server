package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ConfigRoutes(api fiber.Router) {
	group := api.Group("/config")
	handler := handlers.NewConfigHandler(a.db.ConfigStore())

	group.Get("/", handler.GetConfigs)
	group.Get("/:id", handler.GetConfig)
	group.Post("/", handler.CreateConfig)
	group.Patch("/", handler.UpdateConfig)
	group.Delete("/:id", handler.DeleteConfig)
}
