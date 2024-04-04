package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDRIVER             string        `mapstructure:"DB_DRIVER"`
	DBSOURCE             string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	MIGRATION_PATH       string        `mapstructure:"MIGRATION_PATH"`
	REDIS_ADDRESS        string        `mapstructure:"REDIS_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
