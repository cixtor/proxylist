package main

import (
	"flag"
	"fmt"
)

const service = "http://proxylist.hidemyass.com/"

var exportTable = flag.Bool("table", true, "Print proxy servers in an ASCII table")
var exportJSON = flag.Bool("json", false, "Print JSON encoded proxy server list")
var exportCSV = flag.Bool("csv", false, "Print CSV encoded proxy server list")

func main() {
	flag.Usage = func() {
		fmt.Println("ProxyList")
		fmt.Println("http://cixtor.com/")
		fmt.Println("https://github.com/cixtor/proxylist")
		fmt.Println("https://en.wikipedia.org/wiki/Proxy_server")
		fmt.Println(service)
		fmt.Println()
		fmt.Println("In computer networks, a proxy server is a server (a computer system or an\n" +
			"application) that acts as an intermediary for requests from clients seeking\n" +
			"resources from other servers. A client connects to the proxy server, requesting\n" +
			"some service, such as a file, connection, web page, or other resource available\n" +
			"from a different server and the proxy server evaluates the request as a way to\n" +
			"simplify and control its complexity. Proxies were invented to add structure and\n" +
			"encapsulation to distributed systems. Today, most proxies are web proxies,\n" +
			"facilitating access to content on the World Wide Web and providing anonymity.\n")

		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	flag.Parse()

	stream := htmlDocument(service)
	lines := tableCells(stream)

	if *exportCSV {
		printAsCSV(lines)
		return
	}

	if *exportJSON {
		printAsJSON(lines)
		return
	}

	if *exportTable {
		printAsTable(lines)
		return
	}
}
