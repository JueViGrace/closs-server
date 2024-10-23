package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) CompanyRoutes() {
	g := a.Group("/api/company")
	h := handlers.NewCompanyHandler(a.db.CompanyStore())

	// g.Post("/validate")
	g.Get("/", h.GetCompanies)
	g.Get("/:id", h.GetCompany)
	g.Post("/", h.CreateCompany)
	g.Patch("/", h.UpdateCompany)
	g.Delete("/:id", h.DeleteCompany)
}
