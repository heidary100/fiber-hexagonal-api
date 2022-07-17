package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID
	Name      string
	Username  string
	AddedAt   time.Time
	UpdatedAt time.Time
}
