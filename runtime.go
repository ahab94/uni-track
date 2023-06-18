package app

import (
	"context"

	"github.com/ahab94/uni-track/db"

	"github.com/ahab94/uni-track/models"
)

type DatapointStore interface {
	SaveDatapoint(ctx context.Context, dp models.Datapoint) error
	GetDatapointNearest(ctx context.Context, poolID, blockNumber string) (*models.Datapoint, error)
	ListDatapoint(ctx context.Context, filter map[string]interface{}, limit int64) ([]models.Datapoint, error)
}

type Runtime struct {
	datapointStore DatapointStore
}

func NewRuntime() (*Runtime, error) {
	client, err := db.NewMongoClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &Runtime{
		datapointStore: client.Collection("datapoint"),
	}, nil
}

func DefaultRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) DatapointStore() DatapointStore {
	return r.datapointStore
}
