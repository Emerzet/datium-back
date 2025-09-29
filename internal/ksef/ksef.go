package ksef

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

type InitAuthResponse struct {
	Challenge string    `json:"challenge"`
	Timestamp time.Time `json:"timestamp"`
}

//func NewClient(baseURL string, httpClient *http.Client) *Client {}

//func (c *Client) Health(ctx context.Context) error {}

//func (c *Client) InitAuth(ctx context.Context) (string, time.Time, error) {}
