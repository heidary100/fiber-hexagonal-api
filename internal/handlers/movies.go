package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"net/http"
)

func Search(service ports.MoviesService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sr := new(presenter.SearchRequest)

		err := c.QueryParser(sr)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(err))
		}
		titles, err := service.Search(sr.Name)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   titles,
			"err":    err,
		})
	}
}
