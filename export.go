package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printAsTable(lines []string) {
	var count int

	fmt.Println("## |     Alive | Speed | Conn | Anonimity | Protocol + Address + Port     | Country")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, line := range lines {
		count++
		proxy := analyze(line)

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

func printAsJSON(lines []string) {
	var entries []Proxy

	for _, line := range lines {
		proxy := analyze(line)
		entries = append(entries, proxy)
	}

	json.NewEncoder(os.Stdout).Encode(entries)
}
