package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := svc.GetUser(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		res := GetUserResponse{Id: user.Id, Name: user.Name, Age: user.Age}
		return res, nil
	}
}

func MakeCreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		id, err := svc.CreateUser(ctx, req.Name, req.Age)
		if err != nil {
			return nil, err
		}
		res := CreateUserResponse{Id: id}
		return res, nil
	}
}
