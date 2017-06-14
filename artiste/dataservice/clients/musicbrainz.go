package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type MusicbrainzClient struct {
	Client http.Client
	BaseUrl string
}

type Artist struct {
	Name string `json:"name"`
}

func NewMusicBrainzClient(baseUrl string) *MusicbrainzClient {
	return &MusicbrainzClient{
		BaseUrl:baseUrl,
	}
}

func (mc *MusicbrainzClient) GetArtist(artistId string) (Artist, error) {

	resp, err := mc.Client.Get(fmt.Sprintf("%s/ws/2/artist/%s?inc=aliases+releases&fmt=json", mc.BaseUrl, artistId))

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
