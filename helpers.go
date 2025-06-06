package main

import (
	"os"
)

func hasASNDatabase() bool {
	return len(os.Getenv("ASN")) > 0
}

func hasCityDatabase() bool {
	return len(os.Getenv("CITY")) > 0
}

func hasCountryDatabase() bool {
	return len(os.Getenv("COUNTRY")) > 0
}
