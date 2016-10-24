package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func printAsTable(entries ProxyList) {
	fmt.Println("## | Alive      | Speed | Conn | Anonymity | Protocol + Address + Port     | Country")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, proxy := range entries {
		fmt.Printf("%02d | %s | %s | %s | %s | %s | %s\n",
			proxy.Unique,
			padright(proxy.LastUpdate, 10),
			padleft(proxy.Speed, 5),
			padleft(proxy.Connection, 4),
			padright(proxy.Anonymity, 9),
			padright(fullhost(proxy), 29),
			proxy.Country,
		)
	}
}

func printAsJSON(entries ProxyList) {
	json.NewEncoder(os.Stdout).Encode(entries)
}

func printAsCSV(entries ProxyList) {
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
		"Anonymity",
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
			proxy.Anonymity,
		}

		writer.Write(entry)
	}

	writer.Flush()
}
