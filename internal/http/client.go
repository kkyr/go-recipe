package http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	userAgent   = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
	httpTimeout = 15 * time.Second
)

func NewClient() *Client {
	return &Client{
		c: &http.Client{Timeout: httpTimeout},
	}
}

type Client struct {
	c *http.Client
}

func (c *Client) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request failed: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code in response: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return b, nil
}
