package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/handlers"
)

func MovieRouter(app fiber.Router, service ports.MoviesService) {
	app.Get("/movies/search", handlers.Search(service))
}
