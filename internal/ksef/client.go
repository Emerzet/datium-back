package ksef

import (
	"fmt"
	"net/http"
	"strings"
	"time"
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



func NewKsefClient(cfg KsefConfig) *KsefClient {
	// 1.Walidacja BaseURL
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("ksef: BaseURL cannot be empty")

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


	// 4. Zbuduj klienta

	return &KsefClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},

	



	}

}


