package main

import "fmt"

const service = "http://proxylist.hidemyass.com/"

func main() {
	stream := htmlDocument(service)
	lines := tableCells(stream)

	fmt.Printf("%#v\n", lines)
}
