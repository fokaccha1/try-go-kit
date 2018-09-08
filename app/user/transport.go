package user

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeGetUserHttpHandler(svc UserService) http.Handler {
	return httptransport.NewServer(
		MakeGetUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
	)
}

func MakeCreateUserHttpHandler(svc UserService) http.Handler {
	return httptransport.NewServer(
		MakeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
	)
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
