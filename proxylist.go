package main

import (
	"regexp"
)

type Proxy struct {
	LastUpdate string
	Address    string
	Port       string
	Country    string
	Speed      string
	Connection string
	Protocol   string
	Anonimity  string
}

func (p *Proxy) ParseLastUpdate(line string) {
	re := regexp.MustCompile(`<span class="updatets[^>]+>(\S+)`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.LastUpdate = parts[1]
	}
}
