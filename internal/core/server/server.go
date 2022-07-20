package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/routes"
)

type Server struct {
	userHandlers  ports.UserHandler
	filmHandlers  ports.FilmHandler
	musicHandlers ports.MusicHandler
}

func NewServer(userHandlers ports.UserHandler, filmHandlers ports.FilmHandler, musicHandler ports.MusicHandler) *Server {
	return &Server{
		userHandlers:  userHandlers,
		filmHandlers:  filmHandlers,
		musicHandlers: musicHandler,
	}
}

func (s *Server) Initialize(port string) error {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.UserRouter(api, s.userHandlers)
	routes.FilmRouter(api, s.filmHandlers)
	routes.MusicRouter(api, s.musicHandlers)

	return app.Listen(port)
}
