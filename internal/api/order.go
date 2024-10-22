package api

func (a *api) OrderRoutes() {
	orderGroup := a.Group("/api/orders")

	orderGroup.Get("/")
	orderGroup.Get("/:id")
	orderGroup.Post("/")
	orderGroup.Patch("/:id")
	orderGroup.Delete("/:id")
}
