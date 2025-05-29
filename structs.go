package main

type Download struct {
	Folder   string
	Format   string
	Type     string
	CDN      string
	Licenses []string
}

type DataToLoad struct {
	Download Download
	Path     string
	Version  int
}

type IpCity struct {
	IpRangeStart string
	IpRangeEnd   string
	CountryCode  string
	State1       string
	State2       string
	City         string
	Postcode     string
	Latitude     float64
	Longitude    float64
	Timezone     string
	IpVersion    int
	DbVersion    int
}
