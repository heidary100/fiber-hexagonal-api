package services

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/pkg/spotify"
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
	var d []domain.Music
	result, err := spotify.Search(q)
	if err != nil {
		return d, err
	}
	for _, eachResult := range result.Tracks.Tracks {
		music := domain.Music{Name: eachResult.Name}
		d = append(d, music)
	}

	return d, nil
}
