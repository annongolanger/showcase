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

type Artist struct{
	Name string `json:"name"`
}

func(mc *MusicbrainzClient) GetArtist(artistId string) (Artist, error) {

	resp, err := mc.Client.Get(fmt.Sprintf("http://musicbrainz.org/ws/2/artist/%s?inc=aliases+releases&fmt=json", artistId))

	if err != nil {
		return Artist{}, err
	}

	if resp.StatusCode != 200 {
		return Artist{}, errors.New(fmt.Sprintf("Request Error: %s", resp.StatusCode))
	}

	var artistResponse Artist

	err = json.NewDecoder(resp.Body).Decode(&artistResponse)

	if err != nil {
		return Artist{}, errors.New(fmt.Sprintf("Parse Error: %s", err.Error()))
	}

	return artistResponse, nil
}
