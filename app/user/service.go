package user

import (
	"context"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (User, error)
	CreateUser(_ context.Context, name string, age int) (id int, err error)
}

type userService struct {
	uRepo UserRepository
}

func NewService() UserService {
	uRepo := NewUserRepository()
	return &userService{uRepo}
}

func (s *userService) GetUser(ctx context.Context, id int) (User, error) {
	user, err := s.uRepo.Find(ctx, id)
	if err == nil {
		return user, nil
	}
	return User{}, err
}

func (s *userService) CreateUser(ctx context.Context, name string, age int) (id int, err error) {
	id, err = s.uRepo.Store(ctx, name, age)
	return
}
