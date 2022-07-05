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

func MovieErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
