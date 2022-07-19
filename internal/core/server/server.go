package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/routes"
)

type Server struct {
	userHandlers ports.UserHandler
	filmHandlers ports.FilmHandler
}

func NewServer(userHandlers ports.UserHandler, filmHandlers ports.FilmHandler) *Server {
	return &Server{
		userHandlers: userHandlers,
		filmHandlers: filmHandlers,
	}
}

func (s *Server) Initialize(port string) error {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.UserRouter(api, s.userHandlers)
	routes.FilmRouter(api, s.filmHandlers)

	return app.Listen(port)
}
