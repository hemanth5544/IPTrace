package main

import (
	"net/http"
)

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{ "message": "IP-Trace" }`))
}
