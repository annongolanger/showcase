package dataservice

import (
	"context"
	"errors"
	"github.com/benwaine/artistprof/artiste/dataservice/config"
	"github.com/go-kit/kit/endpoint"
)

var ErrNoConfiguredArtists = errors.New("No Configured Artists")

type GetSupportedArtistsResponse struct {
	Artists []Artist `json:"artists"`
	Err     string   `json:"err,omitempty"`
}

//go:generate counterfeiter . GetSupportedArtists
type GetSupportedArtists interface {
	GetSupportedArtists() ([]Artist, error)
}

type SupportedArtistsService struct {
	Config config.ArtisteConfig
}

// GetSupportedArtists returns the supported artists the API consumer can pass into GetArtist
func (a SupportedArtistsService) GetSupportedArtists() ([]Artist, error) {

	if len(a.Config.SupportedArtists) < 1 {
		return []Artist{}, ErrNoConfiguredArtists
	}

	var artists []Artist

	for i := range a.Config.SupportedArtists {
		artists = append(artists, Artist{
			Name: a.Config.SupportedArtists[i].Name,
		})
	}

	return artists, nil
}

func MakeGetSupportedArtistsEndpoint(svc GetSupportedArtists) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		artists, err := svc.GetSupportedArtists()

		if err != nil {
			return GetSupportedArtistsResponse{[]Artist{}, err.Error()}, err
		}

		return GetSupportedArtistsResponse{artists, ""}, nil
	}
}
