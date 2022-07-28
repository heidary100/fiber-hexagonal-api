package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

func FileRouter(app fiber.Router, handler ports.FileHandler) {
	app.Get("/files/find", handler.Find)
}
