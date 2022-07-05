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

	// InsertBook(book *entities.Book) (*entities.Book, error)
	// FetchBooks() (*[]presenter.Book, error)
	// UpdateBook(book *entities.Book) (*entities.Book, error)
	// RemoveBook(ID string) error
}

type MoviesService interface {
	Search(Name string) error
	FetchMovieUrls(Name string) (*[]presenter.Movie, error)
}
