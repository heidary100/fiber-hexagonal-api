package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type postgresRepository struct {
}

func (p postgresRepository) FindAll(ctx context.Context) ([]*domain.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*domain.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Add(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Update(ctx context.Context, uuid uuid.UUID, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() ports.UserRepository {
	return &postgresRepository{}
}
