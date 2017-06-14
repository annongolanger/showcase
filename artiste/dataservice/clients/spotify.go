package clients

import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type SpotifyClient struct {
	client http.Client
}

type SpotifyRelatedArtist struct {
	Name string
}

type relatedArtistResponse struct {
	Artists []relatedArtistArtistResponse `json:"artists"`
}

type relatedArtistArtistResponse struct {
	Name string `json:"name"`
}

func(sc *SpotifyClient) GetRelatedArtists(artistId string) ([]SpotifyRelatedArtist, error)  {

	resp, err := sc.client.Get(fmt.Sprintf("https://api.spotify.com/v1/artists/%s/related-artists", artistId))

	if err != nil {
		return []SpotifyRelatedArtist{}, err
	}

	if resp.StatusCode != 200 {
		return []SpotifyRelatedArtist{}, errors.New("Response Error")
	}

	var relatedArtists relatedArtistResponse

	err = json.NewDecoder(resp.Body).Decode(&relatedArtists)

	if err != nil {
		return []SpotifyRelatedArtist{}, err
	}

	sportifyRelatedArtists := []SpotifyRelatedArtist{}

	for i := range relatedArtists.Artists {

		relatedArtist := SpotifyRelatedArtist{
			Name: relatedArtists.Artists[i].Name,
		}

		sportifyRelatedArtists = append(sportifyRelatedArtists, relatedArtist)
	}

	return sportifyRelatedArtists, nil
}



