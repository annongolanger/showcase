package clients_test

import (
	. "github.com/benwaine/artiste/dataservice/clients"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Musicbrainz", func() {

	var client *MusicbrainzClient

	BeforeEach(func(){
		client = &MusicbrainzClient{}
	})

	Describe("GetArtist", func(){

		var resp MusicBrainsGetArtistResponse
		var err error

		BeforeEach(func(){
			resp, err = client.GetArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
		})

		FIt("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		FIt("should return the artists name", func() {
			Expect(resp.Name).To(Equal("Nirvana"))
		})


	})

})
