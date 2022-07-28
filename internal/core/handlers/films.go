package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"net/http"
)

type FilmHandler struct {
	filmService ports.FilmService
}

var _ ports.FilmHandler = (*FilmHandler)(nil)

func NewFilmHandlers(filmService ports.FilmService) *FilmHandler {
	return &FilmHandler{
		filmService: filmService,
	}
}

func (h *FilmHandler) Search(c *fiber.Ctx) error {
	q := c.Query("q")
	if q == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}
	result, err := h.filmService.Search(q)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.JSON(result)
}
