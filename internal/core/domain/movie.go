package domain

import (
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
