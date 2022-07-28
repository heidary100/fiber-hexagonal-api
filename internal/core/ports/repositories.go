package ports

import (
	"context"
	"github.com/google/uuid"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
)

type UserRepository interface {
	FindAll(context.Context) ([]*domain.UserList, error)
	FindByID(context.Context, uuid.UUID) (*domain.UserList, error)
	Add(context.Context, *domain.User) error
	Update(context.Context, uuid.UUID, *domain.User) error
	Delete(context.Context, uuid.UUID) error
}

type FileRepository interface {
}

type FilmRepository interface {
}

type MusicRepository interface {
}
