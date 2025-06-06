package main

import (
	"fmt"
	"net/http"

	random "github.com/hemanth5544/iptrace/utils"
)

func getHome(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{ "message": "IP-Trace" }`))
}

func getIp(response http.ResponseWriter, request *http.Request) {

	ipString := request.PathValue("ip")
	fmt.Println("PathValue as params: " + ipString)

	jsonBytes, err := fetchIPJson(ipString)
	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonBytes)
}

func getRandomIp(response http.ResponseWriter, request *http.Request) {

	var ipString string
	ipVersion := request.PathValue("ipVersion")
	if ipVersion == "4" {
		ipString = random.RandomIpv4()
	} else {
		ipString = random.RandomIpv6()
	}

	jsonBytes, err := fetchIPJson(ipString)
	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write([]byte(`{ "error": "` + err.Error() + `" }`))
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonBytes)
}
