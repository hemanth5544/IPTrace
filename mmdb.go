package main

import (
	"github.com/maxmind/mmdbwriter"
	"github.com/oschwald/maxminddb-golang"
)

var mmDb = map[string]*maxminddb.Reader{}
var mmDbWriter *mmdbwriter.Tree

func mmdbConnect() {
	mmdbOpenFile("COUNTRY")
	mmdbOpenFile("ASN")
	mmdbOpenFile("CITY")
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


func mmdbOpenFile(dbType string) {}
