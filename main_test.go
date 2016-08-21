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

func TestAnalyzeOne(t *testing.T) {
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
