package clients

import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type MusicbrainzClient struct{
	Client http.Client
}

type MusicBrainsGetArtistResponse struct{
	Name string `json:"name"`
}

func(mc *MusicbrainzClient) GetArtist(artistId string) (MusicBrainsGetArtistResponse, error) {

	resp, err := mc.Client.Get(fmt.Sprintf("http://musicbrainz.org/ws/2/artist/%s?inc=aliases+releases&fmt=json", artistId))

	if err != nil {
		return MusicBrainsGetArtistResponse{}, err
	}

	if resp.StatusCode != 200 {
		return MusicBrainsGetArtistResponse{}, errors.New(fmt.Sprintf("Request Error: %s", resp.StatusCode))
	}

	var artistResponse MusicBrainsGetArtistResponse

	err = json.NewDecoder(resp.Body).Decode(&artistResponse)

	if err != nil {
		return MusicBrainsGetArtistResponse{}, errors.New(fmt.Sprintf("Parse Error: %s", err.Error()))
	}

	return artistResponse, nil
}
