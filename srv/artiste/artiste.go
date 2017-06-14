package main

import (
	"net/http"
	"fmt"
	"flag"
	"github.com/benwaine/artiste/dataservice/config"
	"github.com/benwaine/artiste/dataservice"
	httptransport "github.com/go-kit/kit/transport/http"
	"context"
	"encoding/json"
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

	svc := dataservice.ArtistService{
		Config: config,
	}

	getAllArtistsHandler := httptransport.NewServer(
		dataservice.MakeGetSupportedArtistsEndpoint(svc),
		decodeEmptyBody,
		encodeResponse,
	)

	http.Handle("/GetSupportedArtists", getAllArtistsHandler)

	http.ListenAndServe(":8082", nil)

}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error  {
	return json.NewEncoder(w).Encode(response)
}

func decodeEmptyBody(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
