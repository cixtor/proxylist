package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func printAsTable(entries []Proxy) {
	var count int

	fmt.Println("## | Alive      | Speed | Conn | Anonimity | Protocol + Address + Port     | Country")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, proxy := range entries {
		count++
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

func printAsJSON(entries []Proxy) {
	json.NewEncoder(os.Stdout).Encode(entries)
}

func printAsCSV(entries []Proxy) {
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

	for _, proxy := range entries {
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
