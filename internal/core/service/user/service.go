package usersservice

import (
	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
)

type service struct {
	repository ports.UsersRepository
}

//NewService is used to create a single instance of the service
func NewService(r ports.UsersRepository) ports.UsersService {
	return &service{
		repository: r,
	}
}

//InsertUser is a service layer that helps insert user
func (s *service) InsertUser(user *domain.User) (*domain.User, error) {
	return s.repository.CreateUser(user)
}

//FetchUsers is a service layer that helps fetch all users
func (s *service) FetchUsers() (*[]presenter.User, error) {
	return s.repository.ReadUser()
}

//UpdateUser is a service layer that helps update users
func (s *service) UpdateUser(user *domain.User) (*domain.User, error) {
	return s.repository.UpdateUser(user)
}

//RemoveUser is a service layer that helps remove users
func (s *service) RemoveUser(ID string) error {
	return s.repository.DeleteUser(ID)
}
