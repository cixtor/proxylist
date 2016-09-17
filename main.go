package main

import "flag"

const service = "http://proxylist.hidemyass.com/"

var exportTable = flag.Bool("table", true, "Print proxy servers in an ASCII table")

func main() {
	flag.Parse()

	stream := htmlDocument(service)
	lines := tableCells(stream)

	if *exportTable {
		printAsTable(lines)
		return
	}
}
