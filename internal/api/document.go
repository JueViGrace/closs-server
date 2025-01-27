package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) DocumentRoutes(api fiber.Router) {
	group := api.Group("/documents")
	adminGroup := group.Group("/admin")
	handler := handlers.NewDocumentHandler(a.db.DocumentStore())

	group.Get("/", a.authenticatedHandler(handler.GetDocuments))
	group.Get("/:code", a.authenticatedHandler(handler.GetDocumentByCode))
	adminGroup.Post("/", a.adminAuthMiddleware, handler.CreateDocument)
	adminGroup.Put("/", a.adminAuthMiddleware, handler.UpdateDocument)
}
