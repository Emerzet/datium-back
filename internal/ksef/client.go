package ksef

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"net/url"
)

type KsefConfig struct {
	BaseURL string
	Timeout time.Duration
	Retry int
	CertPath string

}

type KsefClient struct {
	cfg KsefConfig
	httpClient *http.Client
}



func NewKsefClient(cfg KsefConfig) (*KsefClient, error) {
	// 1.Walidacja BaseURL
 
	u, err := url.Parse(cfg.BaseURL)
		if err != nil {
			return nil, fmt.Errorf("ksef: invalid BaseURL: %w", err)

		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return nil, fmt.Errorf("ksef: BaseUrl must use http or https, got %q", u.Scheme)
		}


	// 2. Domyślne wartości
	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}
	if cfg.Retry < 0 {
		cfg.Retry = 0 
	}

// 3.Normalizacja BaseURL
	
	cfg.BaseURL = strings.TrimRight(cfg.BaseURL, "/")


	// 4. Składanie klienta
	httpClient := &http.Client {
		Timeout: cfg.Timeout,
	}


	// Gotowy 

	return  &KsefClient{
		cfg: cfg,
		httpClient: httpClient,

	},nil

	}







