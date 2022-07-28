package services

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/pkg/tmdb"
)

type filmService struct {
	repository ports.FilmRepository
}

func NewFilmService(r ports.FilmRepository) ports.FilmService {
	return &filmService{
		repository: r,
	}
}

func (s *filmService) Search(q string) ([]domain.FilmSearchResult, error) {
	var films []domain.FilmSearchResult
	result, err := tmdb.Search(q)
	if err != nil {
		return films, err
	}
	for _, eachResult := range result.Results {
		film := domain.FilmSearchResult{
			Title:         eachResult.Title,
			OriginalTitle: eachResult.OriginalTitle,
			ReleaseDate:   eachResult.ReleaseDate,
			VoteAverage:   eachResult.VoteAverage,
			VoteCount:     eachResult.VoteCount,
			Popularity:    eachResult.Popularity,
			PosterPath:    eachResult.PosterPath,
			Overview:      eachResult.Overview,
		}
		films = append(films, film)
	}

	return films, nil
}
