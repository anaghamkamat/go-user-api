package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-user-api/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.Create)
	app.Get("/users/:id", h.Get)
	app.Get("/users", h.List)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)
}
