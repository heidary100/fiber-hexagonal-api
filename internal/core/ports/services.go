package ports

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
)

type UsersService interface {
	InsertUser(user *domain.User) (*domain.User, error)
	FetchUsers() (*[]presenter.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	RemoveUser(ID string) error
}

type MoviesService interface {
	Search(Name string) (presenter.MovieSearchResponse, error)
	FetchMovieUrls(Name string, Extensions []string) ([]presenter.FetchUrlResponse, error)
}
