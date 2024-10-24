package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) DocumentRoutes() {
	g := a.Group("/api/documents")
	h := handlers.NewDocumentHandler(a.db.DocumentStore())

	g.Get("/", h.GetDocuments)
	// g.Get("/:id")
	// g.Post("/")
	// g.Patch("/:id")
	// g.Delete("/:id")
}
