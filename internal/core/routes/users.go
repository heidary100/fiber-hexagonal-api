package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

func UserRouter(app fiber.Router, handler ports.UserHandler) {
	app.Get("/users", handler.Get)
	app.Post("/users", handler.Add)
	app.Put("/users", handler.Update)
	app.Delete("/users", handler.Remove)
}
