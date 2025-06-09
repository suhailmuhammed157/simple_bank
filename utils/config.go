package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env                  string        `mapstructure:"ENV"`
	DataSource           string        `mapstructure:"DATA_SOURCE"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisServerAddress   string        `mapstructure:"REDIS_SERVER_ADDRESS"`
	SymmetricKey         string        `mapstructure:"SYMMETRIC_KEY"`
	TokenDuration        time.Duration `mapstructure:"TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailHost            string        `mapstructure:"EMAIL_HOST"`
	EmailPort            int32         `mapstructure:"EMAIL_PORT"`
	EmailUser            string        `mapstructure:"EMAIL_USER"`
	EmailPassword        string        `mapstructure:"EMAIL_PASSWORD"`
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
