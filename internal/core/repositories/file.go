package repositories

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type fileRepository struct {
}

func NewFileRepository() ports.FileRepository {
	return &fileRepository{}
}
