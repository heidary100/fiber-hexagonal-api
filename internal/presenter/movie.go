package presenter

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Source    string             `json:"source"`
	Links     string             `json:"username"`
	Info      string             `json:"info"`
	AddedDate primitive.DateTime `json:"added-date"`
}

type SearchRequest struct {
	Name string `query:"name"`
}

type SearchResponse struct {
	Page         int          `json:"page"`
	TotalPages   int          `json:"total_pages"`
	TotalResults int          `json:"total_results"`
	Results      []TMDBResult `json:"results"`
}

type TMDBResult struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	VoteAverage string `json:"vote_average"`
	ReleaseDate string `json:"release_date"`
	Overview    string `json:"overview"`
}

func MovieErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   err,
		"error":  err.Error(),
	}
}
