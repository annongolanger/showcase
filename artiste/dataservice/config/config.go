package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type ArtistConfig struct {
	Name string
	Mid  string
}

type ArtisteConfig struct {
	SupportedArtists   []ArtistConfig
	MusicBrainzBaseUrl string
	SongKickBaseUrl    string
}

func(ac ArtisteConfig) Supports(name string) (supported bool, id string) {

	for i := range ac.SupportedArtists {
		if ac.SupportedArtists[i].Name == name {
			return true, ac.SupportedArtists[i].Mid
		}
	}

	return false, ""
}

func ParseConfig(configLocation string) (ArtisteConfig, error) {

	jsonConfig, err := ioutil.ReadFile(configLocation)

	if err != nil {
		return ArtisteConfig{}, errors.New(fmt.Sprintf("Unable to open config file: %s ", configLocation))
	}

	jsonReader := bytes.NewReader(jsonConfig)

	var parsedConfig ArtisteConfig

	err = json.NewDecoder(jsonReader).Decode(&parsedConfig)

	if err != nil {
		return ArtisteConfig{}, errors.New(fmt.Sprintf("Unable to parse config file: %s ", configLocation))
	}

	return parsedConfig, nil

}
