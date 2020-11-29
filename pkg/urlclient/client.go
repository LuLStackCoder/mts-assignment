package urlclient

import (
	"context"
	"io/ioutil"
	"net/http"
)

// Client represent client for get data from url
type Client struct {
	// default http-client, because fasthttp-client don't support context canceling
	clientHTTP *http.Client
}

// GetData ...
func (s *Client) GetData(ctx context.Context, url string, cancel context.CancelFunc) (data []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	// listening channel for canceling request
	go func() {
		<-ctx.Done()
		cancel()
	}()
	defer cancel()

	resp, err := s.clientHTTP.Do(req)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = resp.Body.Close()

	return
}

// NewClient ...
func NewClient(clientHTTP *http.Client) *Client {
	return &Client{
		clientHTTP: clientHTTP,
	}
}
