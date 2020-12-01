package urlclient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client represent client for get data from url
type Client struct {
	// default http-client, because fasthttp-client doesn't support context canceling
	clientHTTP *http.Client
}

// GetData ...
func (s *Client) GetData(ctx context.Context, url string) (data []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := s.clientHTTP.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("status code %d for url %s", resp.StatusCode, url)
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)

	return
}

// NewClient ...
func NewClient(clientHTTP *http.Client) *Client {
	return &Client{
		clientHTTP: clientHTTP,
	}
}
