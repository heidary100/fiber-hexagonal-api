package services

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type userService struct {
	repository ports.UserRepository
}

func NewUserService(r ports.UserRepository) ports.UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) Insert(user *domain.User) (*domain.User, error) {
	return s.repository.Create(user)
}

func (s *userService) Fetch() (*[]domain.User, error) {
	return s.repository.Read()
}

func (s *userService) Update(user *domain.User) (*domain.User, error) {
	return s.repository.Update(user)
}

func (s *userService) Remove(ID string) error {
	return s.repository.Delete(ID)
}
