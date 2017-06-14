package dataservice_test

import (
	"github.com/benwaine/artiste/dataservice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/benwaine/artiste/dataservice/config"
	"github.com/go-kit/kit/endpoint"
	"github.com/benwaine/artiste/dataservice/dataservicefakes"
)

var _ = Describe("ArtistService", func() {

	Describe("The GetSubscribedArtists Method", func() {

		var artistService dataservice.ArtistService

		Context("There are configured artists", func() {

			config := config.ArtisteConfig{
				SupportedArtists: []config.ArtistConfig{
					{
						Name: "Foo Fighters",
					},
					{
						Name: "SlipKnot",
					},
				},
			}

			artistService = dataservice.ArtistService{
				Config: config,
			}

			artists, err := artistService.GetSupportedArtists()

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return 2 artists", func() {
				Expect(artists).To(HaveLen(2))
			})

			It("should return an artist list with the correct artists", func() {
				Expect(artists[0].Name).To(Equal("Foo Fighters"))
				Expect(artists[1].Name).To(Equal("SlipKnot"))
			})

		})

		Context("There are noconfigured artists", func() {

			config := config.ArtisteConfig{
				SupportedArtists: []config.ArtistConfig{
				},
			}

			artistService = dataservice.ArtistService{
				Config: config,
			}

			_, err := artistService.GetSupportedArtists()

			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(dataservice.ErrNoConfiguredArtists))
			})
		})

	})

})

var _ = Describe("Artist Endpoint", func() {

	Describe("GetSupportedArtistsEndpoint", func() {

		var getSupportedArtistsEndpoint endpoint.Endpoint
		var service dataservicefakes.FakeGetSupportedArtists
		var ctx dataservicefakes.FakeContext
		var response interface{}
		var err error

		getSupportedArtistsEndpoint = dataservice.MakeGetSupportedArtistsEndpoint(&service)

		Context("Artists are returned with no error", func() {

			service.GetSupportedArtistsReturns([]dataservice.Artist{
				{
					Name: "Test Test",
				},
			}, nil)

			response, err = getSupportedArtistsEndpoint(&ctx, nil)

			It("should call GetSupportedArtists", func() {
				Expect(service.GetSupportedArtistsCallCount()).To(Equal(1))
			})

			It("should not error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should return a response containing the correct artists", func() {
				resp := response.(dataservice.GetSupportedArtistsResponse)
				Expect(resp.Artists[0].Name).To(Equal("Test Test"))
			})
		})

	})

})
