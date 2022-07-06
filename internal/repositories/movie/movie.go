package moviesrepository

import (
	"context"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repository struct {
	Collection *mongo.Collection
}

func (r *repository) CreateMovie(movie *domain.Movie) (*domain.Movie, error) {
	movie.ID = primitive.NewObjectID()
	movie.AddedDate = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r repository) ReadMovie() (*[]presenter.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) DeleteMovie(ID string) error {
	//TODO implement me
	panic("implement me")
}

//NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) ports.MoviesRepository {
	return &repository{
		Collection: collection,
	}
}
