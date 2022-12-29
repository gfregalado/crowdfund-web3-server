package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUri  string `mapstructure:"MONGODB_URI"`
	Port   string `mapstructure:"SERVER_PORT"`
	Origin string `mapstructure:"CLIENT_URL"`

	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenMaxAge      int    `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge     int    `mapstructure:"REFRESH_TOKEN_MAX_AGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
}

func LoadConfig() (config Config, err error) {
	godotenv.Load(".env")

	config.DBUri = os.Getenv("MONGODB_URI")
	config.Port = os.Getenv("SERVER_PORT")
	config.Origin = os.Getenv("CLIENT_URL")
	config.AccessTokenPrivateKey = os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	config.AccessTokenPublicKey = os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")

	config.RefreshTokenPrivateKey = os.Getenv("REFRESH_TOKEN_PRIVATE_KEY")
	config.RefreshTokenPublicKey = os.Getenv("REFRESH_TOKEN_PUBLIC_KEY")

	ACCESS_TOKEN_MAX_AGE := os.Getenv("ACCESS_TOKEN_MAX_AGE")
	config.AccessTokenMaxAge, err = strconv.Atoi(ACCESS_TOKEN_MAX_AGE)

	if err != nil {
		fmt.Println("Error during ACCESS_TOKEN_MAX_AGE conversion")
		return
	}

	REFRESH_TOKEN_MAX_AGE := os.Getenv("REFRESH_TOKEN_MAX_AGE")
	config.AccessTokenMaxAge, err = strconv.Atoi(REFRESH_TOKEN_MAX_AGE)

	if err != nil {
		fmt.Println("Error during REFRESH_TOKEN_MAX_AGE conversion")
		return
	}

	if err != nil {
		return
	}

	return
}
