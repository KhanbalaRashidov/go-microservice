package model

import "go-microservice/metadata/pkg/model"

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata`
}
