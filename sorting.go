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
	return p[i].LastUpdate < p[j].LastUpdate
}

func sortByLastUpdate(entries []Proxy) {
	sort.Sort(ProxyList(entries))
}
