package model

import "go-microservice/gen"

func MetadataToProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}

func MetadataFromProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}
