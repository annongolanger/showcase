package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/benwaine/artistprof/artiste/dataservice/config"

)

var _ = Describe("Config", func() {

	Describe("Supports", func(){

		var config ArtisteConfig


		BeforeEach(func(){
			config = ArtisteConfig{
				SupportedArtists: []ArtistConfig{
					{
						Name: "Mc Jagger",
						Mid: "12345",
					},
				},
			}
		})

		It("Returns true and the artist id when an artist is present in config", func() {
			supported, id := config.Supports("Mc Jagger")
			Expect(supported).To(BeTrue())
			Expect(id).To(Equal("12345"))
		})

		It("Returns false when an artist is no present in config", func() {
			supported, id := config.Supports("No Doubt")
			Expect(supported).To(BeFalse())
			Expect(id).To(BeEmpty())
		})

	})

})
