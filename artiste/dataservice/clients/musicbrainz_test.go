package clients_test

import (
	. "github.com/benwaine/artistprof/artiste/dataservice/clients"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Musicbrainz", func() {

	var client *MusicbrainzClient

	BeforeEach(func() {
		client = &MusicbrainzClient{}
	})

	Describe("GetArtist", func() {

		var resp Artist
		var err error

		Context("A succesfull response is returned", func() {

			BeforeEach(func() {

				mockResponse := ` {
  				"begin_area": {
					"disambiguation": "",
					"id": "a640b45c-c173-49b1-8030-973603e895b5",
					"name": "Aberdeen",
					"sort-name": "Aberdeen"
				},
				"end_area": null,
				"id": "5b11f4ce-a62d-471e-81fc-a69a8278c7da",
				"releases": [
					{
						"packaging-id": null,
						"quality": "normal",
						"barcode": "5018615200826",
						"country": "GB",
						"id": "180abab4-8480-4d1b-9a9c-db3d5e864390",
						"status-id": "4e304316-386d-3409-af2e-78857eec5cfe",
						"title": "Blew",
						"release-events": [
							{
								"date": "1989-12",
								"area": {
									"sort-name": "United Kingdom",
									"iso-3166-1-codes": [
										"GB"
									],
									"disambiguation": "",
									"id": "8a754a16-0027-3a29-b6d7-2b40ea0481ed",
									"name": "United Kingdom"
								}
							}
						],
						"aliases": [

						],
						"date": "1989-12",
						"text-representation": {
							"language": "eng",
							"script": "Latn"
						},
						"status": "Official",
						"disambiguation": "",
						"packaging": null
					}
				],
				"type-id": "e431f5f6-b5d2-343d-8b36-72607fffb74b",
				"ipis": [

				],
				"gender-id": null,
				"isnis": [
					"0000000123486830"
				],
				"life-span": {
					"ended": true,
					"begin": "1988-01",
					"end": "1994-04-05"
				},
				"area": {
					"sort-name": "United States",
					"iso-3166-1-codes": [
						"US"
					],
					"disambiguation": "",
					"id": "489ce91b-6658-3307-9877-795b68554c98",
					"name": "United States"
				},
				"country": "US",
				"type": "Group",
				"name": "Nirvana",
				"gender": null,
				"sort-name": "Nirvana",
				"aliases": [
					{
						"type-id": null,
						"ended": false,
						"sort-name": "Nirvana US",
						"primary": null,
						"end": null,
						"begin": null,
						"type": null,
						"locale": null,
						"name": "Nirvana US"
					}
				],
				"disambiguation": "90s US grunge band"
			}`

				httpmock.RegisterResponder(
					"GET",
					"http://musicbrainz.org/ws/2/artist/5b11f4ce-a62d-471e-81fc-a69a8278c7da?inc=aliases+releases&fmt=json",
					httpmock.NewStringResponder(200, mockResponse))

				resp, err = client.GetArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return the artists name", func() {
				Expect(resp.Name).To(Equal("Nirvana"))
			})
		})

		Context("A unsuccessful response is returned", func() {

			BeforeEach(func() {

				mockResponse := `{"error": "The MusicBrainz web server is currently busy. Please try again later."}`

				httpmock.RegisterResponder(
					"GET",
					"http://musicbrainz.org/ws/2/artist/5b11f4ce-a62d-471e-81fc-a69a8278c7da?inc=aliases+releases&fmt=json",
					httpmock.NewStringResponder(503, mockResponse))

				resp, err = client.GetArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Request Error"))
			})
		})

		Context("A badly formed response is returned", func() {

			BeforeEach(func() {

				mockResponse := `{`

				httpmock.RegisterResponder(
					"GET",
					"http://musicbrainz.org/ws/2/artist/5b11f4ce-a62d-471e-81fc-a69a8278c7da?inc=aliases+releases&fmt=json",
					httpmock.NewStringResponder(200, mockResponse))

				resp, err = client.GetArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Parse Error"))
			})
		})
	})
})
