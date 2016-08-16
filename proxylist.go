package main

import (
	"regexp"
	"strings"
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

func (p *Proxy) InvisibleTags(line string) []string {
	var invisible []string
	var styleStart int = strings.Index(line, "<style>")
	var styleEnd int = strings.Index(line, "</style>")
	var section string = line[(styleStart + 7):styleEnd]

	re := regexp.MustCompile(`\.([0-9a-zA-Z]{4})\{display:(inline|none)\}`)
	parts := re.FindAllStringSubmatch(section, -1)

	for _, data := range parts {
		if data[2] == "none" {
			invisible = append(invisible, data[1])
		}
	}

	return invisible
}
