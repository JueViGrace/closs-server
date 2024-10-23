package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) ConfigRoutes() {
	g := a.Group("/api/config")
	h := handlers.NewConfigHandler(a.db.ConfigStore())

	g.Get("/", h.GetConfigs)
	g.Get("/:id", h.GetConfig)
	g.Post("/", h.CreateConfig)
	g.Patch("/", h.UpdateConfig)
	g.Delete("/:id", h.DeleteConfig)
}
