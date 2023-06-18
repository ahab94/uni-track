package models

type Datapoint struct {
	BlockNumber   string `json:"blockNumber" bson:"_id"`
	PoolID        string `json:"poolID" bson:"poolID"`
	Tick          string `json:"tick" bson:"tick"`
	Token0Balance string `json:"token0Balance" bson:"token0Balance"`
	Token1Balance string `json:"token1Balance" bson:"token1Balance"`
}
