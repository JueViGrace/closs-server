package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CustomerRoutes(api fiber.Router) {
	group := api.Group("/customers")
	adminGroup := group.Group("/admin")

	handler := handlers.NewCustomerHandler(a.db.CustomerStore())

	group.Get("/", a.authenticatedHandler(handler.GetCustomers))
	group.Get("/:code", a.authenticatedHandler(handler.GetCustomerByCode))
	adminGroup.Post("/", a.adminAuthMiddleware, handler.CreateCustomer)
	adminGroup.Put("/", a.adminAuthMiddleware, handler.UpdateCustomer)
}
