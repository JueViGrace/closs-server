package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) ProductRoutes() {
	g := a.Group("/api/products")
	h := handlers.NewProductHandler(a.db.ProductStore())

	g.Get("/", h.GetProducts)
	g.Get("/:id", h.GetProduct)
	g.Post("/", h.CreateProduct)
	g.Patch("/", h.UpdateProduct)
	g.Delete("/:id", h.DeleteProduct)
}
