package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func randomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

// Change to exported function
func RandomIpv4() string {
	numbers := []int{randomNumber(0, 255), randomNumber(0, 255), randomNumber(0, 255), randomNumber(0, 255)}
	var parts []string
	for _, number := range numbers {
		parts = append(parts, strconv.Itoa(number))
	}

	ip := strings.Join(parts, ".")

	// Assuming isIpv4Reserved is defined elsewhere
	if isIpv4Reserved(ip) {
		return RandomIpv4() // Recursive call in case of reserved IP
	}

	return ip
}

func RandomIpv6() string {
	var parts []string

	first := []string{"2001", "2002", "2003", "2400", "2401", "2402", "2403", "2404", "2405", "2406", "2407", "2408",
		"2409", "240a", "2600", "2601", "2602", "2603", "2604", "2605", "2606", "2607", "2608", "2609", "2610", "2620",
		"2800", "2801", "2802", "2803", "2804", "2806", "2a00", "2a01", "2a2", "2a03", "2a04", "2a05", "2a06", "2a07",
		"2a08", "2a09", "2a0a", "2a0b", "2a0c", "2a0d", "2a0e", "2a0f", "2a10", "2a11", "2a12", "2a13", "2a14", "2c0e",
		"2c0f"}

	pick := randomNumber(0, len(first)-1)

	parts = append(parts, first[pick])

	// Generate 7 more parts
	for i := 0; i < 7; i++ {
		parts = append(parts, fmt.Sprintf("%02x", randomNumber(0, 255))+fmt.Sprintf("%02x", randomNumber(0, 255)))
	}

	return strings.Join(parts, ":")
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
