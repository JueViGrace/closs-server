package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) OrderRoutes(api fiber.Router) {
	group := api.Group("/orders")
	handler := handlers.NewOrderHandler(a.db.OrderStore())

	group.Get("/", handler.GetOrders)
	group.Get("/:id", handler.GetOrder)
	group.Post("/", handler.CreateOrder)
	group.Patch("/", handler.UpdateOrder)
	group.Delete("/:id", handler.DeleteOrder)
}
