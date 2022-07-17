package ports

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Read() (*[]domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(ID string) error
}

type FileRepository interface {
	Create(file *domain.File) (*domain.File, error)
	Read() (*[]domain.File, error)
	Update(file *domain.File) (*domain.File, error)
	Delete(ID string) error
}
