package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url    string
	client *http.Client
	apiKey string
}

func New(url, apiKey string) (*Client, error) {
	if url == "" {
		return nil, errors.New("url is empty")
	}

	if apiKey == "" {
		return nil, errors.New("apiKey is empty")
	}

	return &Client{
		url:    url,
		apiKey: apiKey,
		client: &http.Client{},
	}, nil
}

func (c *Client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.url, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", c.apiKey)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, err
}

func (c *Client) post(endpoint string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data to JSON: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.url, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", c.apiKey)
	req.Header.Set("Autharization", c.apiKey)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
