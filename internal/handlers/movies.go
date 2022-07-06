package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"net/http"
	"strings"
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

func GetUrls(service ports.MoviesService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sr := new(presenter.GetUrlsRequest)

		err := c.QueryParser(sr)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MovieErrorResponse(err))
		}
		fmt.Println(sr.Name)
		fmt.Println(sr.Extension)
		gsr, err := service.FetchMovieUrls(sr.Name+` "دانلود رایگان فیلم"`, strings.Split(sr.Extension, ","))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.MovieErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   gsr,
			"err":    err,
		})
	}
}
