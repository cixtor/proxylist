package main

import (
	"regexp"
	"strings"
)

// Proxy defines the JSON object for each proxy.
type Proxy struct {
	Unique     int    `json:"-"`
	Filter     int    `json:"-"`
	LastUpdate string `json:"lastupdate"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Country    string `json:"country"`
	Speed      string `json:"speed"`
	Connection string `json:"connection"`
	Protocol   string `json:"protocol"`
	Anonymity  string `json:"anonymity"`
}

// ParseAddress is used to extract the IPv4 address.
func (p *Proxy) ParseAddress(line string) {
	var end int
	var content string
	var invisible []string

	end = strings.Index(line, "</style>")
	content = line[end+8:]

	// Remove unnecessary empty HTML tags.
	content = strings.Replace(content, "<span></span>", "", -1)

	// Remove explicit invisible HTML tags with inline CSS style.
	re := regexp.MustCompile(`<(span|div) style="display:none">[^<]+<\/(span|div)>`)
	content = re.ReplaceAllString(content, "")

	// Remove implicit invisible HTML tags with CSS classes.
	invisible = p.InvisibleTags(line)
	for _, cssClass := range invisible {
		re = regexp.MustCompile(`<(span|div) class="` + cssClass + `">[^<]+<\/(span|div)>`)
		content = re.ReplaceAllString(content, "")
	}

	// Remove the display inline style from the remaining HTML code.
	re = regexp.MustCompile(`<(span|div) style="display: inline">([^<]+)<\/(span|div)>`)
	content = re.ReplaceAllString(content, "$2")

	// Remove the unnecessary HTML code from visible CSS classes.
	re = regexp.MustCompile(`<(span|div) class="[^"]+">([^<]+)<\/(span|div)>`)
	content = re.ReplaceAllString(content, "$2")

	// Clean final string.
	re = regexp.MustCompile(`([0-9\.]{7,15}).*`)
	content = re.ReplaceAllString(content, "$1")

	if content != "" {
		p.Address = content
	}
}

// InvisibleTags extracts a list of ghost CSS class names.
func (p *Proxy) InvisibleTags(line string) []string {
	var invisible []string
	var section string
	var styleStart int
	var styleEnd int

	styleStart = strings.Index(line, "<style>")
	styleEnd = strings.Index(line, "</style>")
	section = line[(styleStart + 7):styleEnd]

	re := regexp.MustCompile(`\.([0-9a-zA-Z_\-]{4})\{display:(inline|none)\}`)
	parts := re.FindAllStringSubmatch(section, -1)

	for _, data := range parts {
		if data[2] == "none" {
			invisible = append(invisible, data[1])
		}
	}

	return invisible
}

// ParseLastUpdate is used to extract the last update time.
func (p *Proxy) ParseLastUpdate(line string) {
	re := regexp.MustCompile(`<span class="updatets[^>]+>([^<]+)<\/span>`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.LastUpdate = strings.TrimSpace(parts[1])
	}
}

// ParsePort is used to extract the port number.
func (p *Proxy) ParsePort(line string) {
	re := regexp.MustCompile(`<td>(\S+)\s+<\/td>`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Port = parts[1]
	}
}

// ParseCountry is used to extract the country name.
func (p *Proxy) ParseCountry(line string) {
	re := regexp.MustCompile(`\/>([^<]+)<\/span><\/td>`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Country = strings.TrimSpace(parts[1])
	}
}

// ParseSpeed is used to extract the speed in percentage.
func (p *Proxy) ParseSpeed(line string) {
	re := regexp.MustCompile(`style="width: ([0-9]+%);`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Speed = parts[1]
	}
}

// ParseConnection is used to extract the connection speed.
func (p *Proxy) ParseConnection(line string) {
	re := regexp.MustCompile(`style="width: ([0-9]+%);`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Connection = parts[1]
	}
}

// ParseProtocol is used to extract the HTTP transport protocol.
func (p *Proxy) ParseProtocol(line string) {
	re := regexp.MustCompile(`<td>(\S+)\s+<\/td>`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Protocol = parts[1]
	}
}

// ParseAnonymity is used to extract the anonymity percentage.
func (p *Proxy) ParseAnonymity(line string) {
	re := regexp.MustCompile(`<td nowrap>([^>]+)<\/td>`)
	parts := re.FindStringSubmatch(line)

	if len(parts) == 2 {
		p.Anonymity = strings.TrimSpace(parts[1])
	}
}
