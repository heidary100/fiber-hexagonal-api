package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type File struct {
	ID         primitive.ObjectID
	Name       string
	WebPageUrl string
	Url        string
	AddedAt    time.Time
}
