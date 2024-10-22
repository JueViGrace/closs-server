package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) AuthRoutes() {
	authGroup := a.Group("/api/auth")
	authHandler := handlers.NewAuthHandler(a.db.AuthStore())

	authGroup.Post("/signIn", authHandler.SignIn)
	authGroup.Post("/signUp", authHandler.SignUp)
	authGroup.Post("/recover/password")
	authGroup.Post("/recover/username")
}
