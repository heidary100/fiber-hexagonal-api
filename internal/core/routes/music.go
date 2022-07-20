package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

func MusicRouter(app fiber.Router, handler ports.MusicHandler) {
	app.Get("/music/search", handler.Search)
}
