package services

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type fileService struct {
	repository ports.FileRepository
}

func NewFileService(r ports.FileRepository) ports.FileService {
	return &fileService{
		repository: r,
	}
}

func (s *fileService) Insert(file *domain.File) (*domain.File, error) {
	return s.repository.Create(file)
}

func (s *fileService) Fetch() (*[]domain.File, error) {
	return s.repository.Read()
}

func (s *fileService) Update(file *domain.File) (*domain.File, error) {
	return s.repository.Update(file)
}

func (s *fileService) Remove(ID string) error {
	return s.repository.Delete(ID)
}
