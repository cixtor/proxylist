package main

import "fmt"

const service = "http://proxylist.hidemyass.com/"

func main() {
	var count int

	stream := htmlDocument(service)
	lines := tableCells(stream)

	fmt.Println("## |     Alive | Speed | Conn | Anonimity | Protocol + Address + Port     | Country")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, line := range lines {
		count++
		proxy := analyze(line)

		if proxy.Address != "" {
			fmt.Printf("%02d | %s | %s | %s | %s | %s | %s\n",
				count,
				padleft(proxy.LastUpdate, 9),
				padright(proxy.Speed, 5),
				padright(proxy.Connection, 4),
				padright(proxy.Anonimity, 9),
				padright(fullhost(proxy), 29),
				proxy.Country,
			)
		}
	}
}
