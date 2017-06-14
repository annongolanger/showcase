package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/benwaine/artistprof/artiste/dataservice"
	"github.com/benwaine/artistprof/artiste/dataservice/config"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

var configLocation string

func main() {

	flag.StringVar(&configLocation, "config", "config.json", "The location of the JSON config file")
	flag.Parse()

	config, err := config.ParseConfig(configLocation)

	if err != nil {
		panic(fmt.Sprintf("Unable to load configLocation file at %s", configLocation))
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	svc := dataservice.SupportedArtistsService{
		Config: config,
	}

	getAllArtistsHandler := httptransport.NewServer(
		dataservice.MakeGetSupportedArtistsEndpoint(svc),
		decodeEmptyBody,
		encodeResponse,
	)

	http.Handle("/GetSupportedArtists", getAllArtistsHandler)

	getArtistHandler := httptransport.NewServer(
		dataservice.MakeGetArtistEndpoint(),
		dataservice.DecodeGetArtistRequest,
		encodeResponse,
	)

	http.Handle("/GetArtist", getArtistHandler)

	http.ListenAndServe(":8082", nil)

}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeEmptyBody(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
