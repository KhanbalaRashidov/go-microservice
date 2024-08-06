package memory

import (
	"context"
	"go-microservice/rating/internal/repository"
	"go-microservice/rating/pkg/model"
)

type Repository struct {
	data map[model.RecordType]map[model.RecordId][]model.Rating
}

func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordId][]model.Rating{}}
}

func (r *Repository) Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}

	if ratings, ok := r.data[recordType][recordId]; ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}

	return r.data[recordType][recordId], nil
}

func (r *Repository) Put(ctx context.Context, recordID model.RecordId, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordId][]model.
			Rating{}
	}

	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)

	return nil
}
