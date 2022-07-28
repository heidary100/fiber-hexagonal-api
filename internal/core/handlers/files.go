package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"net/http"
)

type FileHandler struct {
	fileService ports.FileService
}

var _ ports.FileHandler = (*FileHandler)(nil)

func NewFileHandlers(fileService ports.FileService) *FileHandler {
	return &FileHandler{
		fileService: fileService,
	}
}

func (h *FileHandler) Find(c *fiber.Ctx) error {
	name := c.Query("name")
	kind := c.Query("kind")
	if name == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}
	result, err := h.fileService.Find(name, kind)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.JSON(result)
}
