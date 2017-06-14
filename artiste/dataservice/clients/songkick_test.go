package clients_test

import (
	. "github.com/benwaine/artistprof/artiste/dataservice/clients"

	. "github.com/onsi/ginkgo"
	"gopkg.in/jarcoal/httpmock.v1"
	. "github.com/onsi/gomega"
)

var _ = Describe("Songkick", func() {

	Describe("GetArtistPerformances", func() {

		var client SongKickClient
		var response []PerformanceEvent
		var err error
		var artistId = "5b11f4ce-a62d-471e-81fc-a69a8278c7da"

		Context("The response is a success", func() {

			var mockResponse = `{
				"resultsPage": {
					"results": {
						"event": [{
							"id": 11129128,
							"type": "Concert",
							"uri": "http://www.songkick.com/concerts/11129128-wild-flag-at-fillmore?utm_source=PARTNER_ID&utm_medium=partner",
							"displayName": "Wild Flag at The Fillmore (April 18, 2012)",
							"start": {
								"time": "20:00:00",
								"date": "2012-04-18",
								"datetime": "2012-04-18T20:00:00-0800"
							},
							"performance": [{
								"artist": {
									"uri": "http://www.songkick.com/artists/29835-wild-flag?utm_source=PARTNER_ID&utm_medium=partner",
									"displayName": "Wild Flag",
									"id": 29835,
									"identifier": []
								},
								"displayName": "Wild Flag",
								"billingIndex": 1,
								"id": 21579303,
								"billing": "headline"
							}],
							"location": {
								"city": "San Francisco, CA, US",
								"lng": -122.4332937,
								"lat": 37.7842398
							},
							"venue": {
								"id": 6239,
								"displayName": "The Fillmore",
								"uri": "http://www.songkick.com/venues/6239-fillmore?utm_source=PARTNER_ID&utm_medium=partner",
								"lng": -122.4332937,
								"lat": 37.7842398,
								"metroArea": {
									"uri": "http://www.songkick.com/metro_areas/26330-us-sf-bay-area?utm_source=PARTNER_ID&utm_medium=partner",
									"displayName": "SF Bay Area",
									"country": {
										"displayName": "US"
									},
									"id": 26330,
									"state": {
										"displayName": "CA"
									}
								}
							},
							"status": "ok",
							"popularity": 0.012763
						}]
					},
					"totalEntries": 24,
					"perPage": 50,
					"page": 1,
					"status": "ok"
				}
			}`

			BeforeEach(func() {

				httpmock.RegisterResponder(
					"GET",
					"http://api.songkick.com/api/3.0/artists/mbid:5b11f4ce-a62d-471e-81fc-a69a8278c7da/calendar.json",
					httpmock.NewStringResponder(200, mockResponse))

				client = SongKickClient{}
				response, err = client.GetArtistPerformances(artistId)
			})

			It("Does not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("Returns an events list", func() {
				Expect(len(response)).To(BeNumerically("==", 1))
			})

		})

		Context("The response is not a 200", func() {

			var mockResponse = `{"resultsPage":{"status":"error","error":{"message":"Invalid or missing apikey"}}}`

			BeforeEach(func() {

				httpmock.RegisterResponder(
					"GET",
					"http://api.songkick.com/api/3.0/artists/mbid:5b11f4ce-a62d-471e-81fc-a69a8278c7da/calendar.json",
					httpmock.NewStringResponder(401, mockResponse))

				client = SongKickClient{}
				response, err = client.GetArtistPerformances(artistId)
			})

			It("returns a request error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Request Error"))
			})
		})

		Context("The response is invalid JSON", func() {

			var mockResponse = `{ lkjsdfsldkfj }`

			BeforeEach(func() {

				httpmock.RegisterResponder(
					"GET",
					"http://api.songkick.com/api/3.0/artists/mbid:5b11f4ce-a62d-471e-81fc-a69a8278c7da/calendar.json",
					httpmock.NewStringResponder(200, mockResponse))

				client = SongKickClient{}
				response, err = client.GetArtistPerformances(artistId)
			})

			It("returns a parse error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Parse Error"))
			})
		})

	})
})
