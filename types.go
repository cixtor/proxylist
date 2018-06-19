package main

// Settings represents the configuration for the proxy.
type Settings struct {
	Protocol       string          `json:"protocol"`
	IP             string          `json:"ip"`
	Type           string          `json:"type"`
	Port           string          `json:"port"`
	Curl           string          `json:"curl"`
	Country        string          `json:"country"`
	IPPort         string          `json:"ipPort"`
	StatusMessage  string          `json:"status_message"`
	AnonymityLevel int             `json:"anonymityLevel"`
	StatusCode     int             `json:"status_code"`
	TsChecked      int64           `json:"tsChecked"`
	Speed          float64         `json:"speed"`
	Get            bool            `json:"get"`
	Post           bool            `json:"post"`
	Cookies        bool            `json:"cookies"`
	Referer        bool            `json:"referer"`
	UserAgent      bool            `json:"user-agent"`
	SupportsHTTPS  bool            `json:"supportsHttps"`
	Websites       map[string]bool `json:"websites"`
}
