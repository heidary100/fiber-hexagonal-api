package moviesservice

import (
	"encoding/json"
	"fmt"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"log"
	"net/http"
	"net/url"
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

func (s *service) Search(name string) (presenter.SearchResponse, error) {
	sr, err := searchInTMDB(name)
	return sr, err
}

func searchInTMDB(name string) (presenter.SearchResponse, error) {
	safeName := url.QueryEscape(name)
	tmdbUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=a8ac2f9446eab16741b3adf87e14cfe9&language=en-US&page=1&include_adult=false&query=%s", safeName)

	var record presenter.SearchResponse
	// Build the request
	req, err := http.NewRequest("GET", tmdbUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return record, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record, err
	}

	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record, nil
}
