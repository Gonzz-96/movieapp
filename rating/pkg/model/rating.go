package model

type UserID string
type RecordID string
type RecordType string
type RatingValue int

const (
	RecordTypeMovie = RecordType("movie")
)

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
