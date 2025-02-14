package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AskRequest struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

type AskResponse struct {
	Text     string   `json:"text"`
	Metadata []string `json:"metadata"`
}

type PromptHistory struct {
	Prompts []PromptReqResponse `json:"prompts"`
}

type PromptReqResponse struct {
	Prompt   string   `json:"prompt"`
	Response string   `json:"response"`
	Metadata []string `json:"metadata"`
}

type AskCoverageRequest struct {
	Text      string `json:"text"`
	FixtureID string `json:"fixture_id"`
	ChatID    string `json:"chat_id"`
}

// Event represents a detailed record of an event.
type Event struct {
	ID             primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string                 `bson:"name" json:"name"`
	Description    string                 `bson:"description,omitempty" json:"description,omitempty"`
	EventType      string                 `bson:"event_type" json:"event_type"`
	EventCategory  string                 `bson:"event_category,omitempty" json:"event_category,omitempty"`
	StartDate      time.Time              `bson:"start_date" json:"start_date"`
	EndDate        time.Time              `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Status         string                 `bson:"status" json:"status"`
	Location       string                 `bson:"location,omitempty" json:"location,omitempty"`
	Participants   []Participant          `bson:"participants,omitempty" json:"participants,omitempty"`
	Tags           []string               `bson:"tags,omitempty" json:"tags,omitempty"`
	LastUpdated    time.Time              `bson:"last_updated" json:"last_updated"`
	AdditionalInfo map[string]interface{} `bson:"additional_info,omitempty" json:"additional_info,omitempty"`
}

// Participant represents an individual or entity participating in an event.
type Participant struct {
	ID       primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string                 `bson:"name" json:"name"`
	Role     string                 `bson:"role,omitempty" json:"role,omitempty"`
	Metadata map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"`
}

type GeneratedContentType string

const (
	GeneratedContentTypeTweet                      GeneratedContentType = "tweet"
	GeneratedContentTypeArticle                    GeneratedContentType = "article"
	GeneratedContentTypePodcastArchiveSeriesScript GeneratedContentType = "podcastArchiveSeriesScript"
)

type GenerateRequest struct {
	Prompt string               `json:"prompt"`
	Lang   string               `json:"lang"`
	Type   GeneratedContentType `json:"type"`
}

// Validate checks if the GenerateRequest has a valid Type
func (req GenerateRequest) Validate() error {
	switch req.Type {
	case GeneratedContentTypeTweet, GeneratedContentTypeArticle, GeneratedContentTypePodcastArchiveSeriesScript:
		return nil
	default:
		return fmt.Errorf("invalid content type: %s", req.Type)
	}
}

type GeneratedPodcastArchiveSeries struct {
	Title        string    `json:"title"`
	Script       string    `json:"script"`
	Tags         []string  `json:"tags"`
	Date         time.Time `json:"date"`
	ShortSummary string    `json:"short_summary"`
}

type GeneratedTweet struct {
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Date time.Time `json:"date"`
}

type GeneratedArticle struct {
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

// GenerateResponse represents a response containing generated content
type GenerateResponse struct {
	Type                 GeneratedContentType           `json:"type"`
	Tweet                *GeneratedTweet                `json:"tweet,omitempty"`
	Article              *GeneratedArticle              `json:"article,omitempty"`
	PodcastArchiveSeries *GeneratedPodcastArchiveSeries `json:"podcastArchiveSeries,omitempty"`
}
