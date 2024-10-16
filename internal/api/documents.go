package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) DocumentRoutes() {
	docGroup := a.Group("/api/documents")

	docHandler := handlers.NewDocumentHandler(a.db.DocumentStore())

	docGroup.Get("/", docHandler.GetDocuments)
	docGroup.Get("/:code", docHandler.GetDocumentsByCode)
	docGroup.Get("/all", docHandler.GetDocumentsWithLines)
	docGroup.Get("/all/:code", docHandler.GetDocumentsWithLinesByCode)
}
