package presenter

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Source    string             `json:"source"`
	Domain    string             `json:"domain"`
	Urls      []string           `json:"urls"`
	Info      MovieInfo          `json:"info"`
	AddedDate time.Time          `json:"added-date"`
}

type MovieInfo struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
}

type SearchRequest struct {
	Name string `query:"name"`
}

type GetUrlsRequest struct {
	Name      string `query:"name"`
	Extension string `query:"ext"`
}

type MovieSearchResponse struct {
	Page         int          `json:"page"`
	TotalPages   int          `json:"total_pages"`
	TotalResults int          `json:"total_results"`
	Results      []TMDBResult `json:"results"`
}

type GoogleSearchResponse struct {
	CurrentPage  int                         `json:"currentPage"`
	Keyword      string                      `json:"keyword"`
	TotalResults int                         `json:"totalResults"`
	Organic      []OrganicGoogleSearchResult `json:"organic"`
}

type OrganicGoogleSearchResult struct {
	Title  string `json:"title"`
	Domain string `json:"domain"`
	Url    string `json:"url"`
}

type TMDBResult struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	PosterPath  string  `json:"poster_path"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
}

func MovieErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   err,
		"error":  err.Error(),
	}
}
