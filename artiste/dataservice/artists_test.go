package dataservice_test

import (
	"github.com/benwaine/artistprof/artiste/dataservice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/benwaine/artistprof/artiste/dataservice/config"
	"github.com/go-kit/kit/endpoint"
	"github.com/benwaine/artistprof/artiste/dataservice/dataservicefakes"
	"errors"
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

		BeforeEach(func() {
			service = dataservicefakes.FakeGetSupportedArtists{}
			ctx = dataservicefakes.FakeContext{}
			getSupportedArtistsEndpoint = dataservice.MakeGetSupportedArtistsEndpoint(&service)
		})

		Describe("The endpoint", func() {

			Context("When artists are returned with no error", func() {

				BeforeEach(func() {
					artists := []dataservice.Artist{{Name: "Test Test", }}
					service.GetSupportedArtistsReturns(artists, nil)
					response, err = getSupportedArtistsEndpoint(&ctx, nil)
				})

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

			Context("An Error is returned", func() {

				BeforeEach(func() {
					service.GetSupportedArtistsReturns([]dataservice.Artist{}, errors.New("Test Error"))
					response, err = getSupportedArtistsEndpoint(&ctx, nil)
				})

				It("should return the error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("Test Error"))
				})
			})
		})
	})
})
