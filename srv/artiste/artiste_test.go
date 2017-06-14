package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

var _ = Describe("Artiste", func() {

	var client *http.Client = &http.Client{}

	Describe("The application is compiled", func() {

		It("Passes the healthcheck", func() {
			resp, err := client.Get("http://localhost:8082/health")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			bodyStr, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			Expect(string(bodyStr)).To(ContainSubstring("OK"))
		})
	})

	Describe("The /GetAllArtists method", func() {

		var response *http.Response
		var body string
		var err error
		var respBytes []byte

		BeforeEach(func() {

			response, err = http.Post("http://localhost:8082/GetSupportedArtists", "application/json", nil)

			if err != nil {
				Fail("Unable to fetch artists resource")
			}

			respBytes, err = ioutil.ReadAll(response.Body)

			if err != nil {
				Fail("Error parsing /GetSupportedArtistsgo response body")
			}

			body = string(respBytes)

		})

		It("should return 200 OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("should return three artists", func() {

			type artist struct {
				Name string
			}

			type artistsResponse struct {
				Artists []artist
			}

			var parsedResp artistsResponse

			err := json.NewDecoder(bytes.NewReader(respBytes)).Decode(&parsedResp)

			Expect(err).NotTo(HaveOccurred())
			Expect(parsedResp.Artists).To(HaveLen(3))
			Expect(parsedResp.Artists[0].Name).To(Equal("Jimmy Eat World"))
			Expect(parsedResp.Artists[1].Name).To(Equal("Nirvana"))
			Expect(parsedResp.Artists[2].Name).To(Equal("New Found Glory"))
		})
	})
})
