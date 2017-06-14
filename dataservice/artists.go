package dataservice

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/benwaine/artiste/dataservice/config"
	"errors"
)

var ErrNoConfiguredArtists = errors.New("No Configured Artists")

type GetSupportedArtistsResponse struct {
	Artists []Artist `json:"artists"`
	Err string `json:"err,omitempty"`
}

type Artist struct {
	Name string `json:"name"`
}

//go:generate counterfeiter . GetSupportedArtists
type GetSupportedArtists interface{
	GetSupportedArtists() ([]Artist, error)
}

type ArtistService struct {
	Config config.ArtisteConfig
}

// GetSupportedArtists returns the supported artists the API consumer can pass into GetArtist
func(a ArtistService) GetSupportedArtists() ([]Artist, error) {

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

	return func (ctx context.Context, request interface{}) (interface{}, error) {

		artists, err := svc.GetSupportedArtists()

		if err != nil {
			return GetSupportedArtistsResponse{[]Artist{}, err.Error()}, nil
		}

		return GetSupportedArtistsResponse{artists, ""}, nil
	}
}