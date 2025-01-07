package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CustomerRoutes(api fiber.Router) {
	group := api.Group("/customers")
	handler := handlers.NewCustomerHandler(a.db.CustomerStore())

	group.Get("/", handler.GetCustomers)
	group.Get("/:id", handler.GetCustomer)
	group.Post("/", handler.CreateCustomer)
	group.Patch("/", handler.UpdateCustomer)
	group.Delete("/:id", handler.DeleteCustomer)
}
