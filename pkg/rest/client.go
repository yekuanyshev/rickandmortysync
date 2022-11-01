package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

func (c *Client) Get(url string, result interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %s, code %d", response.Status, response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}
