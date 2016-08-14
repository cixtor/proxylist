package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

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
	var line string
	var lines []string

	scanner := bufio.NewScanner(stream)

	for scanner.Scan() {
		line = scanner.Text()

		if len(line) > 50 {
			lines = append(lines, line)
		}
	}

	return lines
}
