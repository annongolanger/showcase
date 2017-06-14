package dataservice

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/benwaine/artistprof/artiste/dataservice/clients"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

var ErrInvalidJSON = errors.New("Invalid Request JSON")
var ErrArtistUnavailable = errors.New("Artist Unavailable")
var ErrArtistNotSupported = errors.New("Artist Unsupported")

//go:generate counterfeiter . ArtistGetter
type ArtistGetter interface {
	GetArtist(artistId string) (clients.Artist, error)
}

//go:generate counterfeiter . ArtistPerformanceGetter
type ArtistPerformanceGetter interface {
	GetArtistPerformances(artistId string) ([]clients.PerformanceEvent, error)
}

//go:generate counterfeiter . Supported
type Supported interface {
	Supported(name string) (supported bool, id string)
}

type GetArtistService struct {
	ArtistGetter      ArtistGetter
	PerformanceGetter ArtistPerformanceGetter
	Config            Supported
}

func (s *GetArtistService) GetArtistData(name string) (Artist, error) {

	// Get the ID From config
	supported, id := s.Config.Supported(name)

	if !supported {
		return Artist{}, ErrArtistNotSupported
	}

	// Call the Artist Performance Getter
	artist, err := s.ArtistGetter.GetArtist(id)

	if err != nil {
		return Artist{}, ErrArtistUnavailable
	}

	performances, err := s.PerformanceGetter.GetArtistPerformances(id)

	if err != nil {
		return Artist{}, ErrArtistUnavailable
	}

	artistResponse := Artist{
		Name:         artist.Name,
		Performances: make([]ArtistPerformance, len(performances)),
	}

	for i := range performances {
		artistResponse.Performances[i] = ArtistPerformance{Name: performances[i].Name}
	}

	return artistResponse, nil

	// $$

}

type GetArtistRequest struct {
	Name string `json:"name"`
}

type GetArtistResponse struct {
	Artist Artist
	Error  string
}

func MakeGetArtistEndpoint(service *GetArtistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		getArtistRequest := request.(GetArtistRequest)

		artist, err := service.GetArtistData(getArtistRequest.Name)

		if err != nil {
			return GetArtistResponse{}, err
		}

		return GetArtistResponse{Artist: artist}, nil
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
