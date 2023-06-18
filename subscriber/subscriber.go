package subscriber

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"

	"github.com/ahab94/uni-track/clients"
	"github.com/ahab94/uni-track/config"
	"github.com/ahab94/uni-track/models"
)

const (
	defaultTickSpace = 100
)

type PoolClient interface {
	GetTickData(context.Context, *big.Int, *big.Int) (*big.Int, error)
}

type TokenClient interface {
	GetTokenBalance(context.Context, common.Address, *big.Int) (*big.Int, error)
}

type WSClient interface {
	Subscribe(context.Context, []common.Address) (*clients.Subscription, error)
}

type Subscriber struct {
	contractAddress string
	token0Address   string
	token1Address   string
	poolClient      PoolClient
	tokenClient     TokenClient
	wsClient        WSClient
}

func NewPoolSubscriber(poolCli PoolClient, tokenCli TokenClient,
	wsClient WSClient,
) *Subscriber {
	return &Subscriber{
		wsClient:        wsClient,
		poolClient:      poolCli,
		tokenClient:     tokenCli,
		contractAddress: viper.GetString(config.UniSwapPoolAddress),
		token0Address:   viper.GetString(config.UniSwapToken0Address),
		token1Address:   viper.GetString(config.UniSwapToken1Address),
	}
}

func (u *Subscriber) Subscribe(ctx context.Context, receiver chan<- models.Datapoint) error {
	sub, err := u.wsClient.Subscribe(ctx, []common.Address{common.HexToAddress(u.contractAddress)})
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Sub.Err():
			return err
		case data := <-sub.LogChan:
			bigBlockNumber := big.NewInt(int64(data.BlockNumber))

			tick, err := u.poolClient.GetTickData(ctx, big.NewInt(defaultTickSpace), big.NewInt(int64(data.BlockNumber)))
			if err != nil {
				return err
			}

			token0, err := u.tokenClient.GetTokenBalance(ctx,
				common.HexToAddress(u.token0Address), big.NewInt(int64(data.BlockNumber)))
			if err != nil {
				return err
			}

			token1, err := u.tokenClient.GetTokenBalance(ctx,
				common.HexToAddress(u.token1Address), big.NewInt(int64(data.BlockNumber)))
			if err != nil {
				return err
			}

			receiver <- models.Datapoint{
				Tick:          tick.String(),
				Token0Balance: token0.String(),
				Token1Balance: token1.String(),
				BlockNumber:   bigBlockNumber.String(),
				PoolID:        u.contractAddress,
			}
		}
	}
}
