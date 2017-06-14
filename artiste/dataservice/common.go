package dataservice

type Artist struct {
	Name string `json:"name"`
	Performances []ArtistPerformance
}

type ArtistPerformance struct {
	Name string `json:"name"`
}


