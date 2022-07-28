package ports

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
)

type UserService interface {
	Insert(user *domain.User) (*domain.User, error)
	Fetch() (*[]domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Remove(ID string) error
}

type FileService interface {
	Insert(file *domain.File) (*domain.File, error)
	Fetch() (*[]domain.File, error)
	Update(file *domain.File) (*domain.File, error)
	Remove(ID string) error

	Find(name string, kind string) ([]domain.File, error)
}

type FilmService interface {
	Search(q string) ([]domain.FilmSearchResult, error)
}

type MusicService interface {
	Search(q string) ([]domain.Music, error)
}
