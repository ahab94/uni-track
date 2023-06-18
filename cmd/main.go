package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	UniTrack "github.com/ahab94/uni-track"
	pool "github.com/ahab94/uni-track/api"
	"github.com/ahab94/uni-track/clients"
	"github.com/ahab94/uni-track/config"
	"github.com/ahab94/uni-track/models"
	"github.com/ahab94/uni-track/service"
	"github.com/ahab94/uni-track/subscriber"
)

const (
	defaultHTTPPort = 8080
)

func main() {
	// Initialize configuration
	config.Init()

	// setup app runtime
	runtime, err := UniTrack.NewRuntime()
	if err != nil {
		panic(err)
	}

	// Setup Pool Subscriber
	poolSubscriber := setupPoolSubscriber()

	// Start Pool Subscription
	recv := make(chan models.Datapoint)
	go poolSubscriber.Subscribe(context.Background(), recv) //nolint:errcheck

	// Process received datapoints
	go processDatapoint(recv, service.NewUniTrackService(runtime))

	// run http server
	router := chi.NewRouter()
	api := pool.NewAPI(runtime).NewRouter(router)

	server := newHTTPServer(defaultHTTPPort)

	if err := httpSrvRun(context.Background(), server(api)); err != nil {
		return
	}
}

// setupPoolSubscriber initializes the Pool Subscriber and returns it
func setupPoolSubscriber() *subscriber.Subscriber {
	poolClient := createClient(clients.PoolContractABI)
	tokenClient := createClient(clients.TokenContractABI)
	wsClient := createWebSocketClient()

	return subscriber.NewPoolSubscriber(poolClient, tokenClient, wsClient)
}

// createClient initializes and returns a new UniSwapV3Client
func createClient(abi string) *clients.UniSwapV3Client {
	client, err := clients.NewUniSwapV3Client(viper.GetString(config.EtherNodeURL), abi,
		viper.GetString(config.UniSwapPoolAddress))
	if err != nil {
		panic(err)
	}

	return client
}

// createWebSocketClient initializes and returns a new UniSwapV3PoolWebSocketClient
func createWebSocketClient() *clients.UniSwapV3PoolWebSocketClient {
	wsClient, err := clients.NewUniSwapV3PoolWebSocketClient(viper.GetString(config.EtherNodeWSURL))
	if err != nil {
		panic(err)
	}

	return wsClient
}

// processDatapoint processes received datapoint
func processDatapoint(recv chan models.Datapoint, trackService *service.Service) {
	var prevBlock, currentBlock int
	for dp := range recv {
		currentBlock, _ = strconv.Atoi(dp.BlockNumber)
		if prevBlock != 0 && currentBlock != 0 && currentBlock <= prevBlock {
			logrus.Warnf("Block %d already processed\n", currentBlock)

			continue
		}

		if err := trackService.SaveDatapoint(context.Background(), dp); err != nil {
			logrus.Errorf("save error %+v\n", err)

			continue
		}

		logrus.Infof("Block %d successfully processed\n", currentBlock)

		prevBlock = currentBlock
	}
}

func newHTTPServer(port uint) func(handler http.Handler) *http.Server {
	return func(handler http.Handler) *http.Server {
		addr := fmt.Sprintf(":%d", port)

		srv := http.Server{
			Addr:    addr,
			Handler: handler,
		}

		return &srv
	}
}

func httpSrvRun(_ context.Context, srv *http.Server) error {
	logrus.Infof("Listening HTTP on %s...", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server failed to serve: %w", err)
	}

	return nil
}
