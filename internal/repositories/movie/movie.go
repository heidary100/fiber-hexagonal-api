package moviesrepository

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	Collection *mongo.Collection
}

func (r repository) CreateMovie(user *domain.Movie) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
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
