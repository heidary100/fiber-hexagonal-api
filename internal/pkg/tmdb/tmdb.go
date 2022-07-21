package tmdb

import (
	"encoding/json"
	"fmt"
	"github.com/heidary100/fiber-hexagonal-api/internal/pkg/crawler"
	"net/url"
)

type SearchResponse struct {
	Page         int      `json:"page"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
	Results      []Result `json:"results"`
}

type Result struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	PosterPath  string  `json:"poster_path"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
}

func Search(q string) (SearchResponse, error) {
	var record SearchResponse
	safeName := url.QueryEscape(q)
	endPoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=a8ac2f9446eab16741b3adf87e14cfe9&language=en-US&page=1&include_adult=false&query=%s", safeName)
	responseBody, err := crawler.GetResponseBody(endPoint, crawler.GetHttpClient())
	if err != nil {
		return record, err
	}

	if err := json.NewDecoder(responseBody).Decode(&record); err != nil {
		return record, err
	}

	return record, nil
}
