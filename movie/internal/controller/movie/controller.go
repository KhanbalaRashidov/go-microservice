package movie

import (
	"context"
	"errors"
	metadataModel "go-microservice/metadata/pkg/model"
	"go-microservice/movie/internal/gateway"
	"go-microservice/movie/pkg/model"
	ratingModel "go-microservice/rating/pkg/model"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	Get(ctx context.Context, recordId ratingModel.RecordId, recordType ratingModel.RecordType) (float64, error)
	Put(ctx context.Context, recordId ratingModel.RecordId, recordType ratingModel.RecordType, rating *ratingModel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadataModel.Metadata, error)
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.Get(ctx, ratingModel.RecordId(id), ratingModel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {

	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
