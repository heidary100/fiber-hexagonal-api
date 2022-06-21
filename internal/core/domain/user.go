package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
}

// DeleteRequest struct is used to parse Delete Reqeusts for users
type DeleteRequest struct {
	ID string `json:"id"`
}
