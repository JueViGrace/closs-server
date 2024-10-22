package api

func (a *api) SalesmanRoutes() {
	salesmanGroup := a.Group("/api/salesman")

	salesmanGroup.Get("/")
	salesmanGroup.Get("/:id")
	salesmanGroup.Post("/")
	salesmanGroup.Patch("/:id")
	salesmanGroup.Delete("/:id")
}
