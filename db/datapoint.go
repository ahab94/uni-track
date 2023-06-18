package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahab94/uni-track/models"
)

func (c *Collection) SaveDatapoint(ctx context.Context, dp models.Datapoint) error {
	if _, err := c.collection.InsertOne(ctx, dp); err != nil {
		return errors.Wrap(err, "failed to save datapoint")
	}

	return nil
}

func (c *Collection) GetDatapointNearest(ctx context.Context, poolID, blockNumber string) (*models.Datapoint, error) {
	var dp models.Datapoint

	pipeline := []bson.D{
		{
			{
				"$match",
				bson.D{
					{"poolID", poolID},
				},
			},
		},
	}

	if blockNumber != "" {
		matchStage := bson.D{
			{
				"$match",
				bson.D{
					{"_id", bson.D{{"$lte", blockNumber}}},
				},
			},
		}
		pipeline = append(pipeline, matchStage)
	}

	sortStage := bson.D{
		{
			"$sort",
			bson.D{{"_id", -1}},
		},
	}
	pipeline = append(pipeline, sortStage)

	cursor, err := c.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute aggregation")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		if err := cursor.Decode(&dp); err != nil {
			return nil, errors.Wrap(err, "failed to decode datapoint")
		}
		break
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return &dp, nil
}

func (c *Collection) ListDatapoint(ctx context.Context, filter map[string]interface{}, limit int64) ([]models.Datapoint, error) {
	var dps []models.Datapoint

	cur, err := c.collection.Find(ctx, filter, &options.FindOptions{Limit: &limit, Sort: bson.D{{"_id", -1}}})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list datapoint")
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var dp models.Datapoint

		if err := cur.Decode(&dp); err != nil {
			return nil, errors.Wrap(err, "failed to decode datapoint")
		}

		dps = append(dps, dp)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate datapoint cursor")
	}

	return dps, nil
}
