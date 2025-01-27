package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	group := api.Group("/users")
	handler := handlers.NewUserHandler(a.db.UserStore())

	group.Get("/", a.adminAuthMiddleware, handler.GetUsers)
	group.Get("/me", a.authenticatedHandler(handler.GetUserById))
}
