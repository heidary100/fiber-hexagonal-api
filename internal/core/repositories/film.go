package repositories

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type filmRepository struct {
}

func NewFilmRepository() ports.MusicRepository {
	return &musicRepository{}
}
