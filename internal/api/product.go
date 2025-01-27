package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ProductRoutes(api fiber.Router) {
	group := api.Group("/products")
	adminGroup := api.Group("/admin")
	handler := handlers.NewProductHandler(a.db.ProductStore())

	adminGroup.Get("/", a.adminAuthMiddleware, handler.GetProducts)
	group.Get("/", a.adminAuthMiddleware, handler.GetExistingProducts)
	adminGroup.Get("/:code", a.adminAuthMiddleware, handler.GetProductByCode)
	group.Get("/:code", a.adminAuthMiddleware, handler.GetExistingProductByCode)
	adminGroup.Post("/", handler.CreateProduct)
	adminGroup.Patch("/", handler.UpdateProduct)
	adminGroup.Delete("/:code", handler.DeleteProduct)
}
