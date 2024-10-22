package api

func (a *api) CompanyRoutes() {
	companyGroup := a.Group("/api/company")

	companyGroup.Post("/validate")
	companyGroup.Get("/")
	companyGroup.Get("/:id")
	companyGroup.Post("/")
	companyGroup.Patch("/:id")
	companyGroup.Delete("/:id")
}
