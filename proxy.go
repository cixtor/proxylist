package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Proxy represents the core of the library.
type Proxy struct {
	success int
	service string
	failure []error
	entries []Settings
}

// NewProxy returns an instance of Proxy.
func NewProxy(url string) *Proxy {
	p := new(Proxy)
	p.service = url
	p.failure = make([]error, 0)
	p.entries = make([]Settings, 0)
	return p
}

// Execute returns a list with N proxy settings.
func (p *Proxy) Execute(n int) error {
	fails := make(chan error, n)
	queue := make(chan Settings, n)

	for i := 0; i < n; i++ {
		go p.Fetch(fails, queue)
	}

	var fail error
	var item Settings

	for i := 0; i < n; i++ {
		fail = <-fails
		item = <-queue
		p.failure = append(p.failure, fail)
		p.entries = append(p.entries, item)
		if item.Curl != "" {
			p.success++
		}
	}

	var msg string

	for _, err := range p.failure {
		if err != nil {
			msg += "\xe2\x80\xa2\x20" + err.Error() + "\n"
		}
	}

	if msg == "" {
		return nil
	}

	return errors.New(msg)
}

// Fetch queries a web API service to get one proxy.
func (p *Proxy) Fetch(fails chan error, queue chan Settings) {
	client := &http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("GET", p.service, nil)

	if err != nil {
		fails <- err
		queue <- Settings{}
		return
	}

	req.Header.Set("Host", "gimmeproxy.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-us")

	resp, err := client.Do(req)

	if err != nil {
		fails <- err
		queue <- Settings{}
		return
	}

	defer resp.Body.Close()

	var v Settings

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		fails <- err
		queue <- Settings{}
		return
	}

	if v.StatusCode == 429 {
		fails <- errors.New(v.StatusMessage)
		queue <- Settings{}
		return
	}

	fails <- nil
	queue <- v
}

// Export writes the list of proxies into W in JSON format.
func (p *Proxy) Export(w io.Writer) {
	if err := json.NewEncoder(w).Encode(p.entries); err != nil {
		log.Println("json.decode", err)
	}
}

// Print writes the list of proxies into W in Tabular format.
func (p *Proxy) Print(w io.Writer) {
	var entry []string

	data := [][]string{}

	for _, item := range p.entries {
		if item.Curl == "" {
			continue
		}

		entry = []string{}

		t := time.Unix(item.TsChecked, 0)

		entry = append(entry, item.Country)
		entry = append(entry, item.Curl)
		entry = append(entry, fmt.Sprintf("%.2f", item.Speed))
		entry = append(entry, time.Since(t).String())

		if item.Get {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		if item.Post {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		if item.Cookies {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		if item.Referer {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		if item.UserAgent {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		if item.AnonymityLevel == 1 {
			entry = append(entry, "\x1b[48;5;010m░░\x1b[0m")
		} else {
			entry = append(entry, "\x1b[48;5;009m░░\x1b[0m")
		}

		data = append(data, entry)
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{
		"Country",
		"cURL",
		"Speed",
		"Uptime",
		"G", //            get: bool - supports GET requests
		"P", //           post: bool - supports POST requests
		"C", //        cookies: bool - supports cookies
		"R", //        referer: bool - supports 'referer' header
		"U", //      userAgent: bool - supports 'user-agent' header
		"A", // anonymityLevel:  int - 1:anonymous, 0:notanonymous
	})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	if len(data) > 0 {
		fmt.Fprintln(w, "G - supports GET requests")
		fmt.Fprintln(w, "P - supports POST requests")
		fmt.Fprintln(w, "C - supports cookies")
		fmt.Fprintln(w, "R - supports 'referer' header")
		fmt.Fprintln(w, "U - supports 'user-agent' header")
		fmt.Fprintln(w, "A - 1:anonymous, 0:notanonymous")
	}
}

// Sort re-orders the list of proxies by a column.
func (p *Proxy) Sort(column string) {
	for idx, item := range p.entries {
		switch column {
		case "port":
			p.entries[idx].Filter = item.Port
		case "speed":
			p.entries[idx].Filter = fmt.Sprintf("%.2f", item.Speed*100)
		case "country":
			p.entries[idx].Filter = item.Country
		case "protocol":
			p.entries[idx].Filter = item.Protocol
		case "uptime":
			p.entries[idx].Filter = fmt.Sprintf("%d", item.TsChecked)
		}
	}

	sort.Sort(byFilter(p.entries))
}
