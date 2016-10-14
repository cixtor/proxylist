package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func printAsTable(lines []string) {
	var count int

	fmt.Println("## | Alive      | Speed | Conn | Anonimity | Protocol + Address + Port     | Country")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, line := range lines {
		count++
		proxy := analyze(line)

		fmt.Printf("%02d | %s | %s | %s | %s | %s | %s\n",
			count,
			padright(proxy.LastUpdate, 10),
			padleft(proxy.Speed, 5),
			padleft(proxy.Connection, 4),
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

func printAsCSV(lines []string) {
	var proxy Proxy
	var entry []string

	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{
		"LastUpdate",
		"Address",
		"Port",
		"Country",
		"Speed",
		"Connection",
		"Protocol",
		"Anonimity",
	})

	for _, line := range lines {
		proxy = analyze(line)
		entry = []string{
			proxy.LastUpdate,
			proxy.Address,
			proxy.Port,
			proxy.Country,
			proxy.Speed,
			proxy.Connection,
			proxy.Protocol,
			proxy.Anonimity,
		}

		writer.Write(entry)
	}

	writer.Flush()
}
