package ports

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
)

type UsersRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUser() (*[]presenter.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(ID string) error
}

type MoviesRepository interface {
	CreateMovie(user *domain.Movie) (*domain.Movie, error)
	ReadMovie() (*[]presenter.Movie, error)
	DeleteMovie(ID string) error
}
