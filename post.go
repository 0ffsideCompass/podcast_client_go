package client

import (
	"github.com/0ffsideCompass/models"
)

const (
	createArchieveSeries = "/create/series/archieve"
)

// CreateArchieveSeries creates a new series
func (c *Client) CreateArchieveSeries(prompt string) ([]byte, error) {
	req := models.PodcastArchiveRequests{
		Prompt: prompt,
	}
	return c.post(createArchieveSeries, req)
}
