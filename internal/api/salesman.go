package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) SalesmanRoutes() {
	g := a.Group("/api/salesman")
    h := handlers.NewSalesmanHandler(a.db.SalesmanStore())

	g.Get("/", h.GetSalesmen)
	g.Get("/:id", h.GetSalesman)
	g.Post("/", h.CreateSalesman)
	g.Patch("/", h.UpdateSalesman)
	g.Delete("/:id", h.DeleteSalesman)
}
