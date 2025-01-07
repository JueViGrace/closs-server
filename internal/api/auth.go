package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) AuthRoutes(api fiber.Router) {
	group := api.Group("/auth")
	handler := handlers.NewAuthHandler(a.db.AuthStore())

	group.Post("/signIn", handler.SignIn)
	group.Post("/signUp", handler.SignUp)
	group.Post("/recover/password", handler.RecoverPassword)
}
