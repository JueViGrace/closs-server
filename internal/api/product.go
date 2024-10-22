package api

func (a *api) ProductRoutes() {
	productGroup := a.Group("/api/products")

	productGroup.Get("/")
	productGroup.Get("/:id")
	productGroup.Post("/")
	productGroup.Patch("/:id")
	productGroup.Delete("/:id")
}
