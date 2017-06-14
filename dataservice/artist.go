package dataservice

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"net/http"
	"encoding/json"
	"errors"
)

var ErrInvalidJSON = errors.New("Invalid Request JSON")

type GetArtistRequest struct {
	Name string `json:"name"`
}

type GetArtistResponse struct{}

func MakeGetArtistEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return GetArtistResponse{}, nil
	}
}

func DecodeGetArtistRequest(_ context.Context, r *http.Request) (interface{}, error) {

	defer r.Body.Close()

	request := GetArtistRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return request, ErrInvalidJSON
	}

	return request, err
}
