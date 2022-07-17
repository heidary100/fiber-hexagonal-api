package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"net/http"
)

type UserHandler struct {
	userService ports.UserService
}

var _ ports.UserHandler = (*UserHandler)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	fetched, err := h.userService.Fetch()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}
	return c.JSON(fetched)
}

func (h *UserHandler) Add(c *fiber.Ctx) error {
	var requestBody domain.User
	err := c.BodyParser(&requestBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	//TODO validation
	if requestBody.Username == "" || requestBody.Name == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}
	result, err := h.userService.Insert(&requestBody)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}
	return c.JSON(result)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	var requestBody domain.User
	err := c.BodyParser(&requestBody)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	result, err := h.userService.Update(&requestBody)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}
	return c.JSON(result)
}

func (h *UserHandler) Remove(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}
	err := h.userService.Remove(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}
	return nil
}
