package clients

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type UniSwapV3PoolWebSocketClient struct {
	client *ethclient.Client
}

type Subscription struct {
	Sub        ethereum.Subscription
	HeaderChan chan *types.Header
	LogChan    chan types.Log
}

func NewUniSwapV3PoolWebSocketClient(url string) (*UniSwapV3PoolWebSocketClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return &UniSwapV3PoolWebSocketClient{
		client: client,
	}, nil
}

func (u *UniSwapV3PoolWebSocketClient) Subscribe(ctx context.Context, addresses []common.Address) (*Subscription, error) {
	headers := make(chan *types.Header)
	logs := make(chan types.Log)

	sub, err := u.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{Addresses: addresses}, logs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe")
	}

	return &Subscription{
		Sub:        sub,
		LogChan:    logs,
		HeaderChan: headers,
	}, nil
}
