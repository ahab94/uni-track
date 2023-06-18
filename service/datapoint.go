package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ahab94/uni-track/models"
)

const defaultLimit = 20

func (s *Service) SaveDatapoint(ctx context.Context, datapoint models.Datapoint) error {
	if err := s.rt.DatapointStore().SaveDatapoint(ctx, datapoint); err != nil {
		return errors.Wrap(err, "failed to save datapoint")
	}

	return nil
}

func (s *Service) GetDatapoint(ctx context.Context, poolID, blockNumber string) (*models.Datapoint, error) {
	dp, err := s.rt.DatapointStore().GetDatapointNearest(ctx, poolID, blockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get datapoint")
	}

	return dp, nil
}

func (s *Service) GetDatapointHistory(ctx context.Context, poolID string) ([]models.Datapoint, error) {
	dp, err := s.rt.DatapointStore().ListDatapoint(ctx, map[string]interface{}{
		"poolID": poolID,
	}, defaultLimit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get datapoint")
	}

	return dp, nil
}
