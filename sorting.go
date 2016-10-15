package main

import "sort"

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
	return proxy.Unique
}
