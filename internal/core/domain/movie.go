package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Source    string             `json:"source"`
	Links     string             `json:"username"`
	Info      string             `json:"info"`
	AddedDate primitive.DateTime `json:"added-date"`
}
