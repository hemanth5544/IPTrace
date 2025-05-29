package main

import (
	"encoding/json"
	"errors"
	"net"
	"time"
)

func fetchIP(ipString string) (*Ip, error) {
	start := time.Now()

	ip := net.ParseIP(ipString)
	if ip == nil || ip.IsPrivate() || ip.IsLoopback() {
		return nil, errors.New("invalid IP address passed (" + ipString + "); private / loopback IP ranges are not processed")
	}

	IpResult := mmdbIp(ip)
	IpResult.Milliseconds = time.Now().Sub(start).Milliseconds()
	IpResult.Microseconds = time.Now().Sub(start).Microseconds()

	return IpResult, nil
}

func fetchIPJson(ipString string) ([]byte, error) {
	ipResult, err := fetchIP(ipString)
	if err != nil {
		return nil, err
	}

	jsonResult, err := json.Marshal(ipResult)
	if err != nil {
		return nil, errors.New("system error")
	}

	return jsonResult, nil
}
