package api

import "github.com/JueViGrace/clo-backend/internal/handlers"

func (a *api) AuthRoutes() {
	g := a.Group("/api/auth")
	h := handlers.NewAuthHandler(a.db.AuthStore())

	g.Post("/signIn", h.SignIn)
	g.Post("/signUp", h.SignUp)
	g.Post("/recover/password", h.RecoverPassword)
}
