package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/maxminddb-golang"
)

var mmDb = map[string]*maxminddb.Reader{}
var mmDbWriter *mmdbwriter.Tree

func mmdbConnect() {
	mmdbOpenFile("COUNTRY")
	mmdbOpenFile("ASN")
	mmdbOpenFile("CITY")
}

func mmdbOpenFile(key string) {
	if len(os.Getenv(key)) > 0 {
		ipVersions := []int{4, 6}
		for _, ipVersion := range ipVersions {
			connectionId := key + "ipv" + strconv.Itoa(ipVersion)
			filePath := "downloads/" + os.Getenv(key) + "-ipv" + strconv.Itoa(ipVersion) + ".mmdb"

			if _, err := os.Stat(filePath); err == nil {
				_, ok := mmDb[connectionId]
				if !ok {
					fmt.Println("Opening MMDB file: " + filePath)
					conn, err := maxminddb.Open(filePath)
					if err != nil {
						panic(err)
					}
					//mmDd is a map of connectionId to maxminddb.Reader it will retivere taht connections which were mde by ip if again the request comes
					// This allows us to have multiple connections for different IP versions
					mmDb[connectionId] = conn
				}
			}
		}
	}
}

func mmdbClose() {
	for connectionId, conn := range mmDb {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
		delete(mmDb, connectionId)
	}
}

func mmdbInitialised(key string) bool {
	connectionId := key + "ipv4"
	_, ok := mmDb[connectionId]

	return ok
}

func mmdbCloseFile(connectionId string, filePath string) {
	conn, ok := mmDb[connectionId]
	if ok {
		fmt.Println("Closing MMDB file: " + filePath)
		err := conn.Close()
		if err != nil {
			panic(err)
		}
		delete(mmDb, connectionId)
	}
}

func mmdbIp(ip net.IP) *Ip {
	ipString := ip.String()
	ipVersion := 4
	if strings.Contains(ipString, ":") {
		ipVersion = 6
	}


}

