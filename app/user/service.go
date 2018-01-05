package user

import (
	"context"
	"errors"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (User, error)
	CreateUser(_ context.Context, user User) (id int, err error)
}

type userService struct{}

func NewService() UserService {
	return &userService{}
}

var (
	users = map[int]User{
		1: User{Id: 1, Name: "Jack", Age: 13},
		2: User{Id: 2, Name: "George", Age: 25},
		3: User{Id: 3, Name: "Philip", Age: 38},
	}
	ErrNotFound = errors.New("Not Found")
)

func (*userService) GetUser(_ context.Context, id int) (User, error) {
	user, ok := users[id]
	if ok {
		return user, nil
	}
	return User{}, ErrNotFound
}

func (*userService) CreateUser(_ context.Context, user User) (id int, err error) {
	id = 4
	return id, nil
}
