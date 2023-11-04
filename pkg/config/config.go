package config

import "github.com/spf13/viper"

type Config struct {
	// Database
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	// Server
	ServerAddress string `mapstructure:"APP_ADDRESS"`

  // Authentication
	TokenSymmetricKey   string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration string `mapstructure:"ACCESS_TOKEN_DURATION"`
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
