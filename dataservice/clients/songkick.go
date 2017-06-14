package clients

import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type SongKickClient struct {
	Client http.Client
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

type SongKickEvent struct {
	Name string
}

func (sc SongKickClient) GetArtistPerformances(artistId string) ([]SongKickEvent, error) {

	response, err := sc.Client.Get(fmt.Sprintf("http://api.songkick.com/api/3.0/artists/mbid:%s/calendar.json", artistId))

	if err != nil {
		return []SongKickEvent{}, errors.New(fmt.Sprintf("Request Error: %s", err.Error()))
	}

	if response.StatusCode != 200 {
		return []SongKickEvent{}, errors.New("Request Error")
	}

	var results getPerformanceResponse
	err = json.NewDecoder(response.Body).Decode(&results)

	if err != nil {
		return []SongKickEvent{}, errors.New(fmt.Sprintf("Parse Error %s", err.Error()))
	}

	list := make([]SongKickEvent, len(results.ResultsPage.Event.Events))

	for i := range results.ResultsPage.Event.Events {
		list[i] = SongKickEvent{
			Name: results.ResultsPage.Event.Events[i].DisplayName,
		}
	}

	return list, nil
}
