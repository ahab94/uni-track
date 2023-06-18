package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	LogLevel = "log.level"

	EtherNodeURL   = "ether.node.url"
	EtherNodeWSURL = "ether.node.ws.url"

	UniSwapPoolAddress   = "uniswap.pool.address" //nolint:gosec
	UniSwapToken0Address = "uniswap.pool.token0"  //nolint:gosec
	UniSwapToken1Address = "uniswap.pool.token1"  //nolint:gosec

	DbName = "db.name"
	DbHost = "db.host"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"
)

func Init() {
	_ = viper.BindEnv(LogLevel, "LOG_LEVEL")

	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	_ = viper.BindEnv(EtherNodeURL, "ETHER_NODE_URL")
	_ = viper.BindEnv(EtherNodeWSURL, "ETHER_NODE_WS_URL")

	_ = viper.BindEnv(UniSwapPoolAddress, "UNISWAP_POOL_ADDRESS")
	_ = viper.BindEnv(UniSwapToken0Address, "UNISWAP_TOKEN0_ADDRESS")
	_ = viper.BindEnv(UniSwapToken1Address, "UNISWAP_TOKEN1_ADDRESS")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// defaults
	viper.SetDefault(LogLevel, "debug")

	viper.SetDefault(DbName, "myDB")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")

	viper.SetDefault(EtherNodeURL, "https://mainnet.infura.io/v3/24b9906200114e29bef2fe1aca011d8c")
	viper.SetDefault(EtherNodeWSURL, "wss://mainnet.infura.io/ws/v3/24b9906200114e29bef2fe1aca011d8c")

	viper.SetDefault(UniSwapPoolAddress, "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640")
	viper.SetDefault(UniSwapToken0Address, "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	viper.SetDefault(UniSwapToken1Address, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
}
