package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) SalesmanRoutes(api fiber.Router) {
	group := api.Group("/salesman")
	handler := handlers.NewSalesmanHandler(a.db.SalesmanStore())

	group.Get("/:id", a.authenticatedHandler(handler.GetSalesman))
	group.Post("/", a.adminAuthMiddleware, handler.CreateSalesman)
	group.Patch("/", a.adminAuthMiddleware, handler.UpdateSalesman)
}
