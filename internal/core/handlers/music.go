package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"net/http"
)

type MusicHandler struct {
	musicService ports.MusicService
}

var _ ports.MusicHandler = (*MusicHandler)(nil)

func NewMusicHandlers(musicService ports.MusicService) *MusicHandler {
	return &MusicHandler{
		musicService: musicService,
	}
}

func (h *MusicHandler) Search(c *fiber.Ctx) error {
	q := c.Query("q")
	if q == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}
	result, err := h.musicService.Search(q)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}
	return c.JSON(result)
}
