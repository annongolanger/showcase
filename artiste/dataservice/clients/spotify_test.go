package clients_test

import (
	. "github.com/benwaine/artistprof/artiste/dataservice/clients"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spotify", func() {

	var client *SpotifyClient

	BeforeEach(func() {
		client = &SpotifyClient{}
	})

	Describe("GetReleatedArtists", func() {

		var artists []SpotifyRelatedArtist
		var err error

		BeforeEach(func() {
			Skip("Skipping Spotify")
			artists, err = client.GetRelatedArtists("3Ayl7mCk0nScecqOzvNp6s")
		})

		It("Should not return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return related artists", func() {
			Expect(len(artists)).Should(BeNumerically(">", 1))
		})

	})

})
