package main

import (
	"sort"
	"strconv"
)

type ProxyList []Proxy

func (p ProxyList) Len() int {
	return len(p)
}

func (p ProxyList) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ProxyList) Less(i int, j int) bool {
	return p[i].Filter < p[j].Filter
}

func sortLines(lines []string, sorting string) ProxyList {
	var count int
	var proxy Proxy
	var entries ProxyList

	for _, line := range lines {
		proxy = analyze(line)
		proxy.Unique = count + 1
		proxy.Filter = sortableData(proxy, sorting)
		entries = append(entries, proxy)
		count += 1
	}

	sort.Sort(ProxyList(entries))

	return entries
}

func sortableData(proxy Proxy, sorting string) int {
	if sorting == "speed" {
		return sortableDataBySpeed(proxy)
	}

	if sorting == "connection" {
		return sortableDataByConnection(proxy)
	}

	if sorting == "anonimity" {
		return sortableDataByAnonimity(proxy)
	}

	if sorting == "protocol" {
		return sortableDataByProtocol(proxy)
	}

	if sorting == "port" {
		return sortableDataByPort(proxy)
	}

	return proxy.Unique
}

func sortableDataBySpeed(proxy Proxy) int {
	speed := proxy.Speed[0 : len(proxy.Speed)-1]
	num, _ := strconv.Atoi(speed)

	return num * -1
}

func sortableDataByConnection(proxy Proxy) int {
	connection := proxy.Connection[0 : len(proxy.Connection)-1]
	num, _ := strconv.Atoi(connection)

	return num * -1
}

func sortableDataByAnonimity(proxy Proxy) int {
	reference := map[string]int{
		"High +KA": 1,
		"High":     2,
		"Medium":   3,
		"Low":      4,
		"None":     5,
	}

	return reference[proxy.Anonimity]
}

func sortableDataByProtocol(proxy Proxy) int {
	reference := map[string]int{
		"socks4/5": 1,
		"HTTPS":    2,
		"HTTP":     3,
	}

	return reference[proxy.Protocol]
}

func sortableDataByPort(proxy Proxy) int {
	num, _ := strconv.Atoi(proxy.Port)

	return num
}
