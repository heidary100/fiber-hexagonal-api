package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/handlers"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service ports.UsersService) {
	app.Get("/users", handlers.GetUsers(service))
	app.Post("/users", handlers.AddUser(service))
	app.Put("/users", handlers.UpdateUser(service))
	app.Delete("/users", handlers.RemoveUser(service))
}
