package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func ReportFailure(t *testing.T, test int, expected Proxy) {
	content := FileContent(test)
	buffer := bytes.NewBuffer(content)
	lines := tableCells(buffer)

	if len(lines) == 0 {
		t.Fatal("Base64 encoded test file was not loaded")
	}

	result := analyze(lines[0])

	if result != expected {
		t.Fatalf("Proxy data for Test#%d is incorrect\n"+
			"-: %#v\n+: %#v\n", /* Result in diff format */
			test,
			expected,
			result)
	}
}

func FileContent(test int) []byte {
	var output string

	fpath := fmt.Sprintf("tests/test-%d.txt", test)
	file, err := os.Open(fpath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output += scanner.Text()
	}

	data, err := base64.StdEncoding.DecodeString(output)

	if err != nil {
		panic(err)
	}

	return data
}

func TestAnalyzeProxy1(t *testing.T) {
	ReportFailure(t, 1, Proxy{
		LastUpdate: "56secs",
		Address:    "218.191.247.51",
		Port:       "80",
		Country:    "Hong Kong",
		Speed:      "92%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy2(t *testing.T) {
	ReportFailure(t, 2, Proxy{
		LastUpdate: "56secs",
		Address:    "201.241.88.63",
		Port:       "80",
		Country:    "Chile",
		Speed:      "83%",
		Connection: "94%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy3(t *testing.T) {
	ReportFailure(t, 3, Proxy{
		LastUpdate: "18mins",
		Address:    "95.213.194.94",
		Port:       "3128",
		Country:    "Russia",
		Speed:      "10%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy4(t *testing.T) {
	ReportFailure(t, 4, Proxy{
		LastUpdate: "18mins",
		Address:    "77.251.233.243",
		Port:       "80",
		Country:    "Netherlands",
		Speed:      "100%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy5(t *testing.T) {
	ReportFailure(t, 5, Proxy{
		LastUpdate: "18mins",
		Address:    "60.21.209.114",
		Port:       "8080",
		Country:    "China",
		Speed:      "34%",
		Connection: "91%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy6(t *testing.T) {
	ReportFailure(t, 6, Proxy{
		LastUpdate: "19mins",
		Address:    "202.148.4.26",
		Port:       "8080",
		Country:    "Indonesia",
		Speed:      "45%",
		Connection: "96%",
		Protocol:   "HTTP",
		Anonimity:  "Low",
	})
}

func TestAnalyzeProxy7(t *testing.T) {
	ReportFailure(t, 7, Proxy{
		LastUpdate: "27mins",
		Address:    "220.225.245.109",
		Port:       "8000",
		Country:    "India",
		Speed:      "11%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy8(t *testing.T) {
	ReportFailure(t, 8, Proxy{
		LastUpdate: "28mins",
		Address:    "124.240.187.80",
		Port:       "80",
		Country:    "China",
		Speed:      "88%",
		Connection: "92%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy9(t *testing.T) {
	ReportFailure(t, 9, Proxy{
		LastUpdate: "28mins",
		Address:    "218.191.247.51",
		Port:       "8380",
		Country:    "Hong Kong",
		Speed:      "92%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy10(t *testing.T) {
	ReportFailure(t, 10, Proxy{
		LastUpdate: "28mins",
		Address:    "91.134.221.52",
		Port:       "80",
		Country:    "Bulgaria",
		Speed:      "10%",
		Connection: "40%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy11(t *testing.T) {
	ReportFailure(t, 11, Proxy{
		LastUpdate: "30mins",
		Address:    "137.135.166.225",
		Port:       "8140",
		Country:    "Canada",
		Speed:      "91%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy12(t *testing.T) {
	ReportFailure(t, 12, Proxy{
		LastUpdate: "30mins",
		Address:    "202.112.31.203",
		Port:       "1080",
		Country:    "China",
		Speed:      "62%",
		Connection: "93%",
		Protocol:   "socks4/5",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy13(t *testing.T) {
	ReportFailure(t, 13, Proxy{
		LastUpdate: "31mins",
		Address:    "119.29.233.254",
		Port:       "1080",
		Country:    "Singapore",
		Speed:      "92%",
		Connection: "96%",
		Protocol:   "socks4/5",
		Anonimity:  "High +KA",
	})
}
