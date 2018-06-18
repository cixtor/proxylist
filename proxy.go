package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Proxy represents the core of the library.
type Proxy struct {
	service string
}

// NewProxy returns an instance of Proxy.
func NewProxy(url string) *Proxy {
	p := new(Proxy)
	p.service = url
	return p
}

// List returns a list with N proxy settings.
func (p *Proxy) List(n int) []Settings {
	list := make([]Settings, 0)
	queue := make(chan Settings, n)

	for i := 0; i < n; i++ {
		go p.Fetch(queue)
	}

	var item Settings

	for i := 0; i < n; i++ {
		item = <-queue
		list = append(list, item)
	}

	return list
}

// Fetch queries a web API service to get one proxy.
func (p *Proxy) Fetch(queue chan Settings) {
	client := &http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("GET", p.service, nil)

	if err != nil {
		log.Println("http.NewRequest", err)
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
		log.Println("client.Do", err)
		queue <- Settings{}
		return
	}

	defer resp.Body.Close()

	var v Settings

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Println("json.decode", err)
		queue <- Settings{}
		return
	}

	queue <- v
}
