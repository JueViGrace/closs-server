package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) CustomerRoutes() {
	g := a.Group("/api/customers")
	h := handlers.NewCustomerHandler(a.db.CustomerStore())

	g.Get("/", h.GetCustomers)
	g.Get("/:id", h.GetCustomer)
	g.Post("/", h.CreateCustomer)
	g.Patch("/", h.UpdateCustomer)
	g.Delete("/:id", h.DeleteCustomer)
}
