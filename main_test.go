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

func TestAnalyzeProxy14(t *testing.T) {
	ReportFailure(t, 14, Proxy{
		LastUpdate: "31mins",
		Address:    "113.255.49.49",
		Port:       "80",
		Country:    "Hong Kong",
		Speed:      "92%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy15(t *testing.T) {
	ReportFailure(t, 15, Proxy{
		LastUpdate: "36mins",
		Address:    "122.96.59.106",
		Port:       "80",
		Country:    "China",
		Speed:      "75%",
		Connection: "92%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy16(t *testing.T) {
	ReportFailure(t, 16, Proxy{
		LastUpdate: "38mins",
		Address:    "124.240.187.78",
		Port:       "81",
		Country:    "China",
		Speed:      "75%",
		Connection: "32%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy17(t *testing.T) {
	ReportFailure(t, 17, Proxy{
		LastUpdate: "44mins",
		Address:    "201.55.46.6",
		Port:       "80",
		Country:    "Brazil",
		Speed:      "92%",
		Connection: "68%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy18(t *testing.T) {
	ReportFailure(t, 18, Proxy{
		LastUpdate: "46mins",
		Address:    "120.25.235.11",
		Port:       "8089",
		Country:    "China",
		Speed:      "20%",
		Connection: "96%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy19(t *testing.T) {
	ReportFailure(t, 19, Proxy{
		LastUpdate: "47mins",
		Address:    "40.113.118.174",
		Port:       "8127",
		Country:    "USA",
		Speed:      "90%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy20(t *testing.T) {
	ReportFailure(t, 20, Proxy{
		LastUpdate: "47mins",
		Address:    "114.217.100.172",
		Port:       "808",
		Country:    "China",
		Speed:      "64%",
		Connection: "45%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy21(t *testing.T) {
	ReportFailure(t, 21, Proxy{
		LastUpdate: "49mins",
		Address:    "61.78.133.143",
		Port:       "8080",
		Country:    "Korea",
		Speed:      "9%",
		Connection: "94%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy22(t *testing.T) {
	ReportFailure(t, 22, Proxy{
		LastUpdate: "56mins",
		Address:    "46.218.85.101",
		Port:       "3129",
		Country:    "France",
		Speed:      "48%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy23(t *testing.T) {
	ReportFailure(t, 23, Proxy{
		LastUpdate: "56mins",
		Address:    "178.62.84.228",
		Port:       "8118",
		Country:    "Russia",
		Speed:      "68%",
		Connection: "40%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy24(t *testing.T) {
	ReportFailure(t, 24, Proxy{
		LastUpdate: "1h",
		Address:    "180.142.128.118",
		Port:       "2226",
		Country:    "China",
		Speed:      "16%",
		Connection: "36%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy25(t *testing.T) {
	ReportFailure(t, 25, Proxy{
		LastUpdate: "1h",
		Address:    "119.18.234.60",
		Port:       "80",
		Country:    "China",
		Speed:      "39%",
		Connection: "93%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy26(t *testing.T) {
	ReportFailure(t, 26, Proxy{
		LastUpdate: "1h 1min",
		Address:    "101.81.68.124",
		Port:       "81",
		Country:    "China",
		Speed:      "67%",
		Connection: "93%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy27(t *testing.T) {
	ReportFailure(t, 27, Proxy{
		LastUpdate: "1h 2mins",
		Address:    "173.161.0.227",
		Port:       "80",
		Country:    "USA",
		Speed:      "7%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High",
	})
}

func TestAnalyzeProxy28(t *testing.T) {
	ReportFailure(t, 28, Proxy{
		LastUpdate: "1h 7mins",
		Address:    "61.238.32.69",
		Port:       "1080",
		Country:    "Hong Kong",
		Speed:      "94%",
		Connection: "94%",
		Protocol:   "socks4/5",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy29(t *testing.T) {
	ReportFailure(t, 29, Proxy{
		LastUpdate: "1h 7mins",
		Address:    "190.147.250.90",
		Port:       "3128",
		Country:    "Colombia",
		Speed:      "67%",
		Connection: "96%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy30(t *testing.T) {
	ReportFailure(t, 30, Proxy{
		LastUpdate: "1h 7mins",
		Address:    "122.195.181.46",
		Port:       "8888",
		Country:    "China",
		Speed:      "34%",
		Connection: "91%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy31(t *testing.T) {
	ReportFailure(t, 31, Proxy{
		LastUpdate: "1h 8mins",
		Address:    "210.101.131.231",
		Port:       "8080",
		Country:    "Korea",
		Speed:      "88%",
		Connection: "93%",
		Protocol:   "HTTP",
		Anonimity:  "Low",
	})
}

func TestAnalyzeProxy32(t *testing.T) {
	ReportFailure(t, 32, Proxy{
		LastUpdate: "1h 14mins",
		Address:    "27.202.234.98",
		Port:       "9999",
		Country:    "China",
		Speed:      "55%",
		Connection: "92%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy33(t *testing.T) {
	ReportFailure(t, 33, Proxy{
		LastUpdate: "1h 16mins",
		Address:    "109.202.102.88",
		Port:       "3128",
		Country:    "Netherlands",
		Speed:      "100%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy34(t *testing.T) {
	ReportFailure(t, 34, Proxy{
		LastUpdate: "1h 16mins",
		Address:    "183.57.17.194",
		Port:       "8081",
		Country:    "China",
		Speed:      "78%",
		Connection: "96%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy35(t *testing.T) {
	ReportFailure(t, 35, Proxy{
		LastUpdate: "1h 19mins",
		Address:    "211.143.146.231",
		Port:       "80",
		Country:    "China",
		Speed:      "86%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy36(t *testing.T) {
	ReportFailure(t, 36, Proxy{
		LastUpdate: "1h 19mins",
		Address:    "223.19.212.30",
		Port:       "80",
		Country:    "Hong Kong",
		Speed:      "91%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy37(t *testing.T) {
	ReportFailure(t, 37, Proxy{
		LastUpdate: "1h 21mins",
		Address:    "203.88.164.99",
		Port:       "8080",
		Country:    "Hong Kong",
		Speed:      "8%",
		Connection: "80%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy38(t *testing.T) {
	ReportFailure(t, 38, Proxy{
		LastUpdate: "1h 29mins",
		Address:    "61.55.135.192",
		Port:       "82",
		Country:    "China",
		Speed:      "26%",
		Connection: "89%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy39(t *testing.T) {
	ReportFailure(t, 39, Proxy{
		LastUpdate: "1h 30mins",
		Address:    "87.255.70.228",
		Port:       "3128",
		Country:    "Moldova",
		Speed:      "44%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy40(t *testing.T) {
	ReportFailure(t, 40, Proxy{
		LastUpdate: "1h 31mins",
		Address:    "52.51.196.63",
		Port:       "8080",
		Country:    "USA",
		Speed:      "35%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy41(t *testing.T) {
	ReportFailure(t, 41, Proxy{
		LastUpdate: "1h 31mins",
		Address:    "64.20.56.132",
		Port:       "8080",
		Country:    "USA",
		Speed:      "75%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy42(t *testing.T) {
	ReportFailure(t, 42, Proxy{
		LastUpdate: "1h 31mins",
		Address:    "92.222.108.124",
		Port:       "3128",
		Country:    "Germany",
		Speed:      "80%",
		Connection: "100%",
		Protocol:   "HTTPS",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy43(t *testing.T) {
	ReportFailure(t, 43, Proxy{
		LastUpdate: "1h 33mins",
		Address:    "124.88.67.30",
		Port:       "83",
		Country:    "China",
		Speed:      "73%",
		Connection: "30%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy44(t *testing.T) {
	ReportFailure(t, 44, Proxy{
		LastUpdate: "1h 33mins",
		Address:    "123.56.74.13",
		Port:       "8080",
		Country:    "China",
		Speed:      "69%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy45(t *testing.T) {
	ReportFailure(t, 45, Proxy{
		LastUpdate: "1h 33mins",
		Address:    "223.16.229.8",
		Port:       "8080",
		Country:    "Hong Kong",
		Speed:      "91%",
		Connection: "95%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy46(t *testing.T) {
	ReportFailure(t, 46, Proxy{
		LastUpdate: "1h 34mins",
		Address:    "124.126.168.65",
		Port:       "8888",
		Country:    "China",
		Speed:      "44%",
		Connection: "36%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy47(t *testing.T) {
	ReportFailure(t, 47, Proxy{
		LastUpdate: "1h 36mins",
		Address:    "14.63.226.198",
		Port:       "80",
		Country:    "Korea",
		Speed:      "30%",
		Connection: "94%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy48(t *testing.T) {
	ReportFailure(t, 48, Proxy{
		LastUpdate: "1h 36mins",
		Address:    "31.186.25.157",
		Port:       "3128",
		Country:    "Netherlands",
		Speed:      "13%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy49(t *testing.T) {
	ReportFailure(t, 49, Proxy{
		LastUpdate: "1h 37mins",
		Address:    "212.47.239.185",
		Port:       "9007",
		Country:    "UK",
		Speed:      "93%",
		Connection: "100%",
		Protocol:   "socks4/5",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy50(t *testing.T) {
	ReportFailure(t, 50, Proxy{
		LastUpdate: "1h 38mins",
		Address:    "146.0.73.14",
		Port:       "80",
		Country:    "Russia",
		Speed:      "33%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy51(t *testing.T) {
	ReportFailure(t, 51, Proxy{
		LastUpdate: "18h 49mins",
		Address:    "119.10.72.34",
		Port:       "80",
		Country:    "China",
		Speed:      "53%",
		Connection: "100%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}

func TestAnalyzeProxy52(t *testing.T) {
	ReportFailure(t, 52, Proxy{
		LastUpdate: "18h 53mins",
		Address:    "113.252.236.96",
		Port:       "8080",
		Country:    "Hong Kong",
		Speed:      "91%",
		Connection: "94%",
		Protocol:   "HTTP",
		Anonimity:  "High +KA",
	})
}
