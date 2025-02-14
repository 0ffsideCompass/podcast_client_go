package models

type PodcastArchiveRequests struct {
	Prompt string `json:"prompt"`
}

type Podcast struct {
	ID          string   `json:"id,omitempty"` // ID will be a string in the request (e.g., "63e01f2e5b648a3d9c8c56e2")
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Script      string   `json:"script,omitempty"`
	AudioURL    string   `json:"audio_url,omitempty"`
	Type        string   `json:"type" validate:"required"`
	Language    string   `json:"language" validate:"required"`
	Tags        []string `json:"tags,omitempty"`
	Date        string   `json:"date,omitempty"`
}
