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
	fileHandlers  ports.FileHandler
}

func NewServer(userHandlers ports.UserHandler, filmHandlers ports.FilmHandler, musicHandlers ports.MusicHandler, fileHandlers ports.FileHandler) *Server {
	return &Server{
		userHandlers:  userHandlers,
		filmHandlers:  filmHandlers,
		musicHandlers: musicHandlers,
		fileHandlers:  fileHandlers,
	}
}

func (s *Server) Initialize(port string) error {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.UserRouter(api, s.userHandlers)
	routes.FilmRouter(api, s.filmHandlers)
	routes.MusicRouter(api, s.musicHandlers)
	routes.FileRouter(api, s.fileHandlers)

	return app.Listen(port)
}
