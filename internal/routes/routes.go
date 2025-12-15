package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-user-api/internal/handler"
	"go-user-api/internal/middleware"
	

)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestID())
app.Use(middleware.RequestLogger())


	app.Post("/users", h.Create)
	app.Get("/users", h.List)
	app.Get("/users/:id", h.Get)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)
}
