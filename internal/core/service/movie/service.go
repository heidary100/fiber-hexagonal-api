package moviesservice

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
)

type service struct {
	repository ports.MoviesRepository
}

func NewService(r ports.MoviesRepository) ports.MoviesService {
	return &service{
		repository: r,
	}
}

func (s *service) FetchMovieUrls(Name string) (*[]presenter.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Search(name string) ([]string, error) {
	return []string{"hello", name}, nil
}
