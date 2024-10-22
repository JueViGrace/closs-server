package api

func (a *api) CustomerRoutes() {
	customerGroup := a.Group("/api/customers")

	customerGroup.Get("/")
	customerGroup.Get("/:id")
	customerGroup.Post("/")
	customerGroup.Patch("/:id")
	customerGroup.Delete("/:id")
}
