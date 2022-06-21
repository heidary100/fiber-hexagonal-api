package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"github.com/pkg/errors"
)

// AddUser is handler/controller which creates Users
func AddUser(service ports.UsersService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if requestBody.Username == "" || requestBody.Name == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(errors.New(
				"Please specify name and username")))
		}
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// UpdateUser is handler/controller which updates data of Users
func UpdateUser(service ports.UsersService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// RemoveUser is handler/controller which removes Users
func RemoveUser(service ports.UsersService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		bookID := requestBody.ID
		err = service.RemoveUser(bookID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

// GetUsers is handler/controller which lists all Users
func GetUsers(service ports.UsersService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(fetched))
	}
}
