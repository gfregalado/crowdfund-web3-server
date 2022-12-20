package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUri  string `mapstructure:"MONGODB_URI"`
	Port   string `mapstructure:"SERVER_PORT"`
	Origin string `mapstructure:"CLIENT_URL"`

	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   int    `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  int    `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int    `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge     int    `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
