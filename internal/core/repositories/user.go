package repositories

import (
	"context"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(conn string) *UserRepository {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}
	return &UserRepository{
		client:     client,
		database:   client.Database("fiber-file-api"),
		collection: client.Database("fiber-file-api").Collection("users"),
	}
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.AddedAt = time.Now()
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Read() (*[]domain.User, error) {
	var users []domain.User
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user domain.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}

func (r *UserRepository) Update(user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
