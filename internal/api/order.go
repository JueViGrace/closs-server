package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) OrderRoutes() {
	g := a.Group("/api/orders")
	h := handlers.NewOrderHandler(a.db.OrderStore())

	g.Get("/", h.GetOrders)
	g.Get("/:id", h.GetOrder)
	g.Post("/", h.CreateOrder)
	g.Patch("/", h.UpdateOrder)
	g.Delete("/:id", h.DeleteOrder)
}
