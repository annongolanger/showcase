package dataservice_test

import (
	. "github.com/benwaine/artistprof/artiste/dataservice"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/go-kit/kit/endpoint"
	"github.com/benwaine/artistprof/artiste/dataservice/dataservicefakes"
	"net/http"
	"bytes"
)

var _ = Describe("Artist", func() {

	Describe("The Endpoint", func() {

		var endpoint endpoint.Endpoint
		var response interface{}
		var request GetArtistRequest
		var err error

		BeforeEach(func() {
			endpoint = MakeGetArtistEndpoint()
			request = GetArtistRequest{}
			ctx := dataservicefakes.FakeContext{}
			response, err = endpoint(&ctx, request)
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
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
