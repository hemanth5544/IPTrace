package main

import (
	"net"
	"os"
	"strings"

	"github.com/seancfoley/ipaddress-go/ipaddr"
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

func getIpVersion(ipString string) int {
	ipVersion := 4
	if strings.Contains(ipString, ":") {
		ipVersion = 6
	}

	return ipVersion
}

func findIPRanges(ipRangeStart string, ipRangeEnd string) []*net.IPNet {
	ipStart := ipaddr.NewIPAddressString(ipRangeStart)
	ipEnd := ipaddr.NewIPAddressString(ipRangeEnd)

	addressStart := ipStart.GetAddress()
	addressEnd := ipEnd.GetAddress()

	ipRange := addressStart.SpanWithRange(addressEnd)
	rangeSlice := ipRange.SpanWithPrefixBlocks()

	var ipNets []*net.IPNet
	for _, val := range rangeSlice {
		_, network, err := net.ParseCIDR(val.String())
		if err != nil {
			panic(err)
		}

		ipNets = append(ipNets, network)
	}

	return ipNets
}
