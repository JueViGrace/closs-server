package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CompanyRoutes(api fiber.Router) {
	group := api.Group("/company")
	adminGroup := group.Group("/admin")

	handler := handlers.NewCompanyHandler(a.db.CompanyStore())

	adminGroup.Get("/", a.adminAuthMiddleware, handler.GetCompanies)
	adminGroup.Get("/:code", a.adminAuthMiddleware, handler.GetCompanyByCode)
	group.Get("/:code", handler.GetExistingCompanyByCode)
	adminGroup.Post("/", a.adminAuthMiddleware, handler.CreateCompany)
	adminGroup.Put("/", a.adminAuthMiddleware, handler.UpdateCompany)
	adminGroup.Delete("/:code", a.adminAuthMiddleware, handler.SoftDeleteCompany)
	adminGroup.Delete("/forever/:code", a.adminAuthMiddleware, handler.DeleteCompany)
}
