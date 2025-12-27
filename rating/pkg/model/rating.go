package model

type UserID string
type RecordID string
type RecordType string
type RatingValue int

type RatingEventType string

const (
	RecordTypeMovie = RecordType("movie")
)

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}

type RatingEvent struct {
	Rating
	ProviderID string          `json:providerId`
	EventType  RatingEventType `json:eventType`
}

const (
	RatingEventTypePut    = RatingEventType("put")
	RatingEventTypeDelete = RatingEventType("delete")
)
