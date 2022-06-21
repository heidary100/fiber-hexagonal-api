package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the presenter object which will be passed in the response by Handler
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
}

// UserSuccessResponse is the singular SuccessResponse that will be passed in the response by
//Handler
func UserSuccessResponse(data *domain.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

// UsersSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
