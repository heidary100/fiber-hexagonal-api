package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

func FilmRouter(app fiber.Router, handler ports.FilmHandler) {
	app.Get("/films/search", handler.Search)
}
