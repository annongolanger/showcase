package config

import (
	"errors"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

type ArtistConfig struct {
	Name string
	Mid string
}

type ArtisteConfig struct {
	SupportedArtists []ArtistConfig
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
