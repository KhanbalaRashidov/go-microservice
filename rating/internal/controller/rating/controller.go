package rating

import (
	"context"
	"errors"
	"go-microservice/rating/internal/repository"
	"go-microservice/rating/pkg/model"
)

var ErrNotFound = errors.New("ratings not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordID model.RecordId, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordId, recordType model.RecordType, rating *model.Rating) error
}

type ratingIngester interface {
	Ingest(ctx context.Context) (chan model.RatingEvent, error)
}

// Controller defines a rating service controller.
type Controller struct {
	repo     ratingRepository
	ingester ratingIngester
}

// New creates a rating service controller.
func New(repo ratingRepository, ingester ratingIngester) *Controller {
	return &Controller{repo, ingester}
}

func (c *Controller) Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordId, recordType)
	if err != nil && err != repository.ErrNotFound {
		return 0, ErrNotFound
	}

	sum := float64(0)
	for _, rating := range ratings {
		sum += float64(rating.Value)
	}

	return sum / float64(len(ratings)), nil
}

func (c *Controller) Put(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, recordId, recordType, rating)
}

func (s *Controller) StartIngestion(ctx context.Context) error {
	ch, err := s.ingester.Ingest(ctx)
	if err != nil {
		return err
	}
	for e := range ch {
		if err := s.Put(ctx, e.RecordId, e.RecordType, &model.Rating{UserId: e.UserId, Value: e.Value}); err != nil {
			return err
		}
	}
	return nil
}
