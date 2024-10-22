package api

func (a *api) UserRoutes() {
	userGroup := a.Group("/api/users")

	userGroup.Get("/")
	userGroup.Get("/:id")
	userGroup.Patch("/:id")
	userGroup.Delete("/:id")
}
