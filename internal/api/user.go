package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) UserRoutes() {
	g := a.Group("/api/users")
	h := handlers.NewUserHandler(a.db.UserStore())

	g.Get("/", h.GetUsers)
	g.Get("/:id", h.GetUser)
	g.Patch("/:id", h.UpdateUser)
	g.Delete("/:id", h.DeleteUser)
}
