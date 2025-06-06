package main

import (
	"os"
	"strings"
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

func isIpv4Reserved(ip string) bool {
	return strings.HasPrefix(ip, "0.") || strings.HasPrefix(ip, "127.") || strings.HasPrefix(ip, "10.") ||
		strings.HasPrefix(ip, "100.64.") || strings.HasPrefix(ip, "169.254.") || strings.HasPrefix(ip, "172.16.") ||
		strings.HasPrefix(ip, "172.17.") || strings.HasPrefix(ip, "172.18.") || strings.HasPrefix(ip, "172.19.") ||
		strings.HasPrefix(ip, "172.20.") || strings.HasPrefix(ip, "172.21.") || strings.HasPrefix(ip, "172.22.") ||
		strings.HasPrefix(ip, "172.23.") || strings.HasPrefix(ip, "172.24.") || strings.HasPrefix(ip, "172.25.") ||
		strings.HasPrefix(ip, "172.26.") || strings.HasPrefix(ip, "172.27.") || strings.HasPrefix(ip, "172.28.") ||
		strings.HasPrefix(ip, "172.29.") || strings.HasPrefix(ip, "172.30.") || strings.HasPrefix(ip, "172.31.") ||
		strings.HasPrefix(ip, "192.0.0.") || strings.HasPrefix(ip, "192.0.2.") || strings.HasPrefix(ip, "192.88.") ||
		strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "198.18.") || strings.HasPrefix(ip, "198.19.") ||
		strings.HasPrefix(ip, "198.51.100") || strings.HasPrefix(ip, "203.0.113") || strings.HasPrefix(ip, "224.") ||
		strings.HasPrefix(ip, "225.") || strings.HasPrefix(ip, "226.") || strings.HasPrefix(ip, "227.") ||
		strings.HasPrefix(ip, "228.") || strings.HasPrefix(ip, "229.") || strings.HasPrefix(ip, "230.") ||
		strings.HasPrefix(ip, "231.") || strings.HasPrefix(ip, "232.") || strings.HasPrefix(ip, "233.") ||
		strings.HasPrefix(ip, "234.") || strings.HasPrefix(ip, "235.") || strings.HasPrefix(ip, "236.") ||
		strings.HasPrefix(ip, "237.") || strings.HasPrefix(ip, "238.") || strings.HasPrefix(ip, "239.") ||
		strings.HasPrefix(ip, "240.") || strings.HasPrefix(ip, "241.") || strings.HasPrefix(ip, "242.") ||
		strings.HasPrefix(ip, "243.") || strings.HasPrefix(ip, "244.") || strings.HasPrefix(ip, "245.") ||
		strings.HasPrefix(ip, "246.") || strings.HasPrefix(ip, "247.") || strings.HasPrefix(ip, "248.") ||
		strings.HasPrefix(ip, "249.") || strings.HasPrefix(ip, "250.") || strings.HasPrefix(ip, "251.") ||
		strings.HasPrefix(ip, "252.") || strings.HasPrefix(ip, "253.") || strings.HasPrefix(ip, "254.") ||
		strings.HasPrefix(ip, "255.")
}
