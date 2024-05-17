package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// Database
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	// Server
	HttpServerAddress string `mapstructure:"HTTP_APP_ADDRESS"`
	RpcServerAddress  string `mapstructure:"RPC_APP_ADDRESS"`
	RunServer         string `mapstructure:"RUN_SERVER"`

	// Authentication
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

	// Redis
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
  RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
