package repositories

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type musicRepository struct {
}

func NewMusicRepository() ports.MusicRepository {
	return &musicRepository{}
}
