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
	AccessTokenMaxAge      int    `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int    `mapstructure:"REFRESH_TOKEN_MAX_AGE"`

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

	viper.SetDefault("MONGODB_URI", "default")
	viper.SetDefault("SERVER_PORT", "default")
	viper.SetDefault("CLIENT_URL", "default")

	viper.SetDefault("ACCESS_TOKEN_PRIVATE_KEY", "default")
	viper.SetDefault("ACCESS_TOKEN_PUBLIC_KEY", "default")
	viper.SetDefault("ACCESS_TOKEN_MAX_AGE", "default")

	viper.SetDefault("REFRESH_TOKEN_PRIVATE_KEY", "default")
	viper.SetDefault("REFRESH_TOKEN_PUBLIC_KEY", "default")
	viper.SetDefault("REFRESH_TOKEN_MAX_AGE", "default")

	if err != nil {
		return
	}

	var c Config

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// ADD START
	envKeysMap := &map[string]interface{}{}
	for k := range *envKeysMap {
		if bindErr := viper.BindEnv(k); bindErr != nil {
			return
		}
	}
	// ADD END

	if err = viper.Unmarshal(&c); err != nil {
		return
	}

	return
}
