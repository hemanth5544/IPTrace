package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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


