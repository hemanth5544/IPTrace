package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func getBenchmark(response http.ResponseWriter, request *http.Request) {

	var ipString string

	ipVersion := request.PathValue("ipVersion")
	times := request.PathValue("times")

	timesInt, err := strconv.Atoi(times)
	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write([]byte(`{ "error": "URL must contain a numeric number of times to run" }`))
		return
	}

	// Don't want this to be part of the benchmark
	var testIps []string
	for i := 0; i < timesInt; i++ {
		if ipVersion == "4" {
			ipString = random.RandomIpv4()
		} else {
			ipString = random.RandomIpv6()
		}

		testIps = append(testIps, ipString)
	}

	start := time.Now()
	for _, ipString := range testIps {
		_, err := fetchIP(ipString)
		if err != nil {
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte(`{ "error": "Error encountered during run (` + ipString + `)" }`))
			return
		}
	}

	ms := int(time.Now().Sub(start).Milliseconds())
	us := int(time.Now().Sub(start).Microseconds())

	msPerOp := ms / timesInt
	usPerOp := us / timesInt

	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{
		"times": ` + strconv.Itoa(timesInt) + `, 
		"ms": ` + strconv.Itoa(ms) + `, 
		"μs": ` + strconv.Itoa(us) + `, 
		"ms_per_op": ` + strconv.Itoa(msPerOp) + `, 
		"μs_per_op": ` + strconv.Itoa(usPerOp) + ` 
	}`))
}