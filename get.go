package client

import (
	"encoding/json"
	"fmt"

	"github.com/0ffsideCompass/models"
)

const (
	searchLanguageEndpoint = "/search/language/%s"
	searchSeriesEndpoint   = "/search/series/%s"
	getPodcast             = "/get/podcasts/%s"
)

// GetLanguages returns a list of languages
func (c *Client) GetLanguages(language string) ([]models.Podcast, error) {
	responseData, err := c.get(fmt.Sprintf(searchLanguageEndpoint, language))
	if err != nil {
		return nil, err
	}

	var data []models.Podcast
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}

	return data, nil
}

// GetSeries returns a list of series
func (c *Client) GetSeries(series string) ([]models.Podcast, error) {
	responseData, err := c.get(fmt.Sprintf(searchSeriesEndpoint, series))
	if err != nil {
		return nil, err
	}

	var data []models.Podcast
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}
	return data, nil
}

// GetPodcast returns a podcast
func (c *Client) GetPodcast(id string) (*models.Podcast, error) {
	responseData, err := c.get(fmt.Sprintf(getPodcast, id))
	if err != nil {
		return nil, err
	}

	var data models.Podcast
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %w", err)
	}
	return &data, nil
}
