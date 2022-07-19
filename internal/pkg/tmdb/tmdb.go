package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
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

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

// question: difference between pointer return values and normal return values
func sendSearchRequest(q string, client *http.Client) (SearchResponse, error) {
	var record SearchResponse
	safeName := url.QueryEscape(q)
	endPoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=a8ac2f9446eab16741b3adf87e14cfe9&language=en-US&page=1&include_adult=false&query=%s", safeName)
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		// todo:
		// log.Fatalf("Error Occurred. %+v", err)
		return record, err
	}

	response, err := client.Do(req)
	if err != nil {
		// log.Fatalf("Error sending request to API endpoint. %+v", err)
		return record, err
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		return record, err
	}

	return record, nil
}

func Search(q string) (SearchResponse, error) {
	c := httpClient()
	return sendSearchRequest(q, c)
}
