package api

func (a *api) ConfigRoutes() {
	configGroup := a.Group("/api/config")

	configGroup.Get("/")
	configGroup.Get("/:id")
	configGroup.Post("/")
	configGroup.Patch("/:id")
	configGroup.Delete("/:id")
}
