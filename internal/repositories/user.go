package usersrepository

import (
	"context"

	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	Collection *mongo.Collection
}

//NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) ports.UsersRepository {
	return &repository{
		Collection: collection,
	}
}

//CreateBook is a mongo repository that helps to create books
func (r *repository) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	// book.CreatedAt = time.Now()
	// book.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//ReadBook is a mongo repository that helps to fetch books
func (r *repository) ReadUser() (*[]presenter.User, error) {
	var users []presenter.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user presenter.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}

//UpdateBook is a mongo repository that helps to update books
func (r *repository) UpdateUser(book *domain.User) (*domain.User, error) {
	// book.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	if err != nil {
		return nil, err
	}
	return book, nil
}

//DeleteBook is a mongo repository that helps to delete books
func (r *repository) DeleteUser(ID string) error {
	bookID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": bookID})
	if err != nil {
		return err
	}
	return nil
}
