package ksef

import (
	"bytes"
	"context"
	"net/http"
)


type Client struct {
	BaseURL string
	HTTP *http.Client
}

type InitAuthResponse struct {
	Challenge string `json:"challenge"`
	Timestamp time.Time `json:"timestamp"`

}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTP: httpClient,

	}

}

func (c *Client) Health(ctx context.Context) error {

	req, _ := http.NewRequestWithContext(ctx,"POST", c.BaseURL+"/auth/init", bytes.NewReader(payload) )


},




func (c *Client) InitAuth(ctx context.Context) (string, time.Time, error) {


}
