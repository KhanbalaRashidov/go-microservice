package model

type RecordId string

type RecordType string

const (
	RecordTypeMovie = RecordType("movie")
)

type UserId string

type RatingValue int

type Rating struct {
	RecordId   RecordId    `json:"recordId"`
	RecordType RecordType  `json:"recordType"`
	UserId     UserId      `json:"userId"`
	Value      RatingValue `json:"value"`
}

type RatingEvent struct {
	UserId     UserId          `json:"userId"`
	RecordId   RecordId        `json:"recordId"`
	RecordType RecordType      `json:"recordType"`
	Value      RatingValue     `json:"value"`
	EventType  RatingEventType `json:"eventType"`
}

type RatingEventType string

const (
	RatingEventTypePut    = "put"
	RatingEventTypeDelete = "delete"
)
