package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CompanyRoutes(api fiber.Router) {
	group := api.Group("/company")
	handler := handlers.NewCompanyHandler(a.db.CompanyStore())

	// g.Post("/validate")
	group.Get("/", handler.GetCompanies)
	group.Get("/:id", handler.GetCompany)
	group.Post("/", handler.CreateCompany)
	group.Patch("/", handler.UpdateCompany)
	group.Delete("/:id", handler.DeleteCompany)
}
