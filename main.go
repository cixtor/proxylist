package main

import "flag"

const service = "http://proxylist.hidemyass.com/"

var exportTable = flag.Bool("table", true, "Print proxy servers in an ASCII table")
var exportJSON = flag.Bool("json", false, "Print JSON encoded proxy server list")

func main() {
	flag.Parse()

	stream := htmlDocument(service)
	lines := tableCells(stream)

	if *exportJSON {
		printAsJSON(lines)
		return
	}

	if *exportTable {
		printAsTable(lines)
		return
	}
}
