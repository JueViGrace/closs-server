package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ConfigRoutes(api fiber.Router) {
	group := api.Group("/config")
	adminGroup := group.Group("/config")

	handler := handlers.NewConfigHandler(a.db.ConfigStore())

	adminGroup.Get("/", a.adminAuthMiddleware, handler.GetConfigs)
	group.Get("/:id", a.authenticatedHandler(handler.GetConfigsByUser))
	adminGroup.Post("/", a.adminAuthMiddleware, handler.CreateConfig)
	adminGroup.Put("/", a.adminAuthMiddleware, handler.UpdateConfig)
	adminGroup.Delete("/", a.adminAuthMiddleware, handler.DeleteConfig)
}
