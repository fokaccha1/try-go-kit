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
		return user, nil
	}
}

func MakeCreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		id, err := svc.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return id, nil
	}
}
