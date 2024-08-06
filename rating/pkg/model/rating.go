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
