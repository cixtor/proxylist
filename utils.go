package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

func analyze(text string) Proxy {
	var proxy Proxy
	text = strings.Replace(text, "</td><", "</td>\n<", -1)
	var lines []string = strings.Split(text, "\n")

	if len(lines) == 8 {
		proxy = cellData(lines)
	}

	return proxy
}

func cellData(lines []string) Proxy {
	var proxy Proxy

	if len(lines) >= 8 {
		proxy.ParseLastUpdate(lines[0])
		proxy.ParseAddress(lines[1])
		proxy.ParsePort(lines[2])
		proxy.ParseCountry(lines[3])
		proxy.ParseSpeed(lines[4])
		proxy.ParseConnection(lines[5])
		proxy.ParseProtocol(lines[6])
		proxy.ParseAnonimity(lines[7])
	}

	return proxy
}

func analyzeLines(lines []string, sorting string) []Proxy {
	var entries []Proxy

	for _, line := range lines {
		entries = append(entries, analyze(line))
	}

	sortByLastUpdate(entries)

	return entries
}

func htmlDocument(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,*/*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")

	if err != nil {
		log.Fatalf("NewRequest: %s", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Client.Do: %s", err)
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	(&buf).ReadFrom(resp.Body)

	return &buf
}

func tableCells(stream io.Reader) []string {
	scanner := bufio.NewScanner(stream)

	var line string
	var cells []string
	var collect bool = false
	var row string

	for scanner.Scan() {
		line = scanner.Text()

		if strings.Contains(line, "<tr class=") {
			collect = true
			continue
		}

		if collect {
			if strings.Contains(line, "</tr>") {
				cells = append(cells, row)
				collect = false
				row = ""
				continue
			}

			row += strings.TrimSpace(line)
		}
	}

	return cells
}

func fullhost(proxy Proxy) string {
	proto := strings.ToLower(proxy.Protocol)

	if proxy.Protocol == "socks4/5" {
		proto = "socks"
	}

	return proto + "://" + proxy.Address + ":" + proxy.Port
}

func padleft(text string, length int) string {
	var total int = len(text)

	if total < length {
		text = strings.Repeat("\x20", length-total) + text
	}

	return text
}

func padright(text string, length int) string {
	var total int = len(text)

	if total < length {
		text += strings.Repeat("\x20", length-total)
	}

	return text
}
