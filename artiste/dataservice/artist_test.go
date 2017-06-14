package dataservice_test

import (
	. "github.com/benwaine/artistprof/artiste/dataservice"

	"bytes"
	"errors"
	"github.com/benwaine/artistprof/artiste/dataservice/clients"
	"github.com/benwaine/artistprof/artiste/dataservice/dataservicefakes"
	"github.com/go-kit/kit/endpoint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Artist", func() {

	Describe("The Endpoint uses the service to get artist data", func() {

		var endpoint endpoint.Endpoint
		var fakeArtistGetter *dataservicefakes.FakeArtistGetter
		var fakePerformanceGetter *dataservicefakes.FakeArtistPerformanceGetter
		var fakeSupported *dataservicefakes.FakeSupported
		var getArtistSrv *GetArtistService
		var response interface{}
		var request GetArtistRequest
		var err error
		var ctx dataservicefakes.FakeContext

		BeforeEach(func() {

			fakeArtistGetter = &dataservicefakes.FakeArtistGetter{}
			fakePerformanceGetter = &dataservicefakes.FakeArtistPerformanceGetter{}
			fakeSupported = &dataservicefakes.FakeSupported{}

			getArtistSrv = &GetArtistService{
				ArtistGetter:      fakeArtistGetter,
				PerformanceGetter: fakePerformanceGetter,
				Config:            fakeSupported,
			}

			endpoint = MakeGetArtistEndpoint(getArtistSrv)
			request = GetArtistRequest{}
			ctx = dataservicefakes.FakeContext{}
		})

		Context("A supported artist is requested", func() {

			BeforeEach(func() {
				fakeSupported.SupportedReturns(true, "123456789")
			})

			Context("Both Artist and Performance Data are returned", func() {

				BeforeEach(func() {
					fakeArtistGetter.GetArtistReturns(clients.Artist{Name: "Sting"}, nil)
					fakePerformanceGetter.GetArtistPerformancesReturns([]clients.PerformanceEvent{{Name: "Live At The Paladium"}}, nil)
					response, err = endpoint(&ctx, request)
				})

				It("should not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("Should return an artist with performance data", func() {
					getArtistResponse := response.(GetArtistResponse)
					Expect(getArtistResponse.Artist.Name).To(Equal("Sting"))
					Expect(len(getArtistResponse.Artist.Performances)).To(BeNumerically("==", 1))
					Expect(getArtistResponse.Artist.Performances[0].Name).To(Equal("Live At The Paladium"))
				})
			})

			Context("Artist Data is not returned", func() {

				BeforeEach(func() {
					fakeArtistGetter.GetArtistReturns(clients.Artist{}, errors.New("An Error"))
					fakePerformanceGetter.GetArtistPerformancesReturns([]clients.PerformanceEvent{{Name: "Live At The Paladium"}}, nil)
					response, err = endpoint(&ctx, request)
				})

				It("should error", func() {
					Expect(err).To(Equal(ErrArtistUnavailable))
				})
			})

			Context("Performance Data is not returned", func() {

				BeforeEach(func() {
					fakeArtistGetter.GetArtistReturns(clients.Artist{Name: "Sting"}, nil)
					fakePerformanceGetter.GetArtistPerformancesReturns([]clients.PerformanceEvent{}, ErrArtistUnavailable)
					response, err = endpoint(&ctx, request)
				})

				It("should error", func() {
					Expect(err).To(Equal(ErrArtistUnavailable))
				})
			})
		})

		Context("An unsupported artist is requested", func() {

			BeforeEach(func() {
				fakeSupported.SupportedReturns(false, "")
				response, err = endpoint(&ctx, request)
			})

			It("should return an error", func() {
				Expect(err).To(Equal(ErrArtistNotSupported))
			})
		})
	})

	Describe("The HTTP Request Decoder", func() {

		var request *http.Request
		var parsedRequest interface{}
		var err error

		Context("valid JSON is submitted", func() {

			BeforeEach(func() {
				reader := bytes.NewReader([]byte(`{ "name": "Jimmy Eat World" }`))
				request, _ = http.NewRequest("POST", "/GetArtist", reader)
				ctx := dataservicefakes.FakeContext{}
				parsedRequest, err = DecodeGetArtistRequest(&ctx, request)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should parse the request into a GetArtistRequest", func() {
				getArtistRequest, success := parsedRequest.(GetArtistRequest)

				Expect(success).To(BeTrue())
				Expect(getArtistRequest.Name).To(Equal("Jimmy Eat World"))
			})
		})

		Context("invalid JSON is submitted", func() {

			BeforeEach(func() {
				reader := bytes.NewReader([]byte(`{ "na`))
				request, _ = http.NewRequest("POST", "/GetArtist", reader)
				ctx := dataservicefakes.FakeContext{}
				parsedRequest, err = DecodeGetArtistRequest(&ctx, request)
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(ErrInvalidJSON))
			})
		})
	})
})
