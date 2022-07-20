package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
)

type userRepository struct {
}

func (r userRepository) FindAll(ctx context.Context) ([]*domain.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*domain.UserList, error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepository) Add(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (r userRepository) Update(ctx context.Context, uuid uuid.UUID, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (r userRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() ports.UserRepository {
	return &userRepository{}
}
