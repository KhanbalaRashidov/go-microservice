package memory

import (
	"context"
	"go-microservice/metadata/internal/repository"
	"go-microservice/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func NewRepository() *Repository {
	return &Repository{data: make(map[string]*model.Metadata)}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {

	r.RLock()
	defer r.RUnlock()

	r.data[id] = metadata

	return nil
}
