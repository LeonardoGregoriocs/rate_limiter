package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel                          string `mapstructure:"LOG_LEVEL"`
	WebServerPort                     int    `mapstructure:"WEB_SERVER_PORT"`
	RedisHost                         string `mapstructure:"REDIS_HOST"`
	RedisPort                         int    `mapstructure:"REDIS_PORT"`
	RedisPassword                     string `mapstructure:"REDIS_PASSWORD"`
	RedisDB                           int    `mapstructure:"REDIS_DB"`
	RateLimiterIPMaxRequests          int    `mapstructure:"RATE_LIMITER_IP_MAX_REQUESTS"`
	RateLimiterTokenMaxRequests       int    `mapstructure:"RATE_LIMITER_TOKEN_MAX_REQUESTS"`
	RateLimiterTimeWindowMilliseconds int    `mapstructure:"RATE_LIMITER_TIME_WINDOW_MILISECONDS"`
}

func GetConfigs(path string) (*Config, error) {
	var cfg *Config

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
