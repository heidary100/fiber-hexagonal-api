package services

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type musicService struct {
	repository ports.MusicRepository
}

func NewMusicService(r ports.MusicRepository) ports.MusicService {
	return &musicService{
		repository: r,
	}
}

func (s *musicService) Search(q string) ([]domain.Music, error) {
	var films []domain.Music
	// TODO spotify
	//result, err := tmdb.Search(q)
	//if err != nil {
	//	return films, err
	//}
	//for _, eachResult := range result.Results {
	//	film := domain.Film{Name: eachResult.Title}
	//	films = append(films, film)
	//}

	return films, nil
}
