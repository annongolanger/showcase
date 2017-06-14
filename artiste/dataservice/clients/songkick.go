package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SongKickClient struct {
	Client http.Client
	BaseUrl string
}

func NewSongKickClient(baseUrl string) *SongKickClient {
	return &SongKickClient{
		BaseUrl: baseUrl,
	}
}

type getPerformanceResponse struct {
	ResultsPage results `json:"resultsPage"`
}

type results struct {
	Event events `json:"results"`
}

type events struct {
	Events []event `json:"event"`
}

type event struct {
	DisplayName string `json:"displayName"`
}

type PerformanceEvent struct {
	Name string
}

func (sc SongKickClient) GetArtistPerformances(artistId string) ([]PerformanceEvent, error) {

	response, err := sc.Client.Get(fmt.Sprintf("%s/api/3.0/artists/mbid:%s/calendar.json", sc.BaseUrl, artistId))

	if err != nil {
		return []PerformanceEvent{}, errors.New(fmt.Sprintf("Request Error: %s", err.Error()))
	}

	if response.StatusCode != 200 {
		return []PerformanceEvent{}, errors.New("Request Error")
	}

	var results getPerformanceResponse
	err = json.NewDecoder(response.Body).Decode(&results)

	if err != nil {
		return []PerformanceEvent{}, errors.New(fmt.Sprintf("Parse Error %s", err.Error()))
	}

	list := make([]PerformanceEvent, len(results.ResultsPage.Event.Events))

	for i := range results.ResultsPage.Event.Events {
		list[i] = PerformanceEvent{
			Name: results.ResultsPage.Event.Events[i].DisplayName,
		}
	}

	return list, nil
}
