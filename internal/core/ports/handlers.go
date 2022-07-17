package ports

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}
