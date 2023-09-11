package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASSWORD"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRATION_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRATION_HOUR"`
	AccessTokenSecret      string `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	RateLimit              int    `mapstructure:"RATE_LIMIT"`
	LoginUsecaseTimeout_MS int    `mapstructure:"LOGIN_USECASE_TIMEOUT_MS"`
	BypassDBPing           bool   `mapstructure:"BYPASS_DB"`
}

func readFromEnv(env *Env) error {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(env)

}

func readFromEnvProd(env *Env) error {
	viper.BindEnv("APP_ENV")
	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("CONTEXT_TIMEOUT")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("JWT_ACCESS_TOKEN_EXPIRATION_HOUR")
	viper.BindEnv("JWT_REFRESH_TOKEN_EXPIRATION_HOUR")
	viper.BindEnv("JWT_ACCESS_TOKEN_SECRET")
	viper.BindEnv("JWT_REFRESH_TOKEN_SECRET")
	viper.BindEnv("RATE_LIMIT")
	viper.BindEnv("LOGIN_USECASE_TIMEOUT_MS")
	viper.BindEnv("BYPASS_DB")

	return viper.Unmarshal(env)
}

func NewEnv() (*Env, error) {
	env := Env{}

	switch env.AppEnv {
	case "development":
		err := readFromEnv(&env)
		if err != nil {
			return nil, err
		}

	case "production":
		err := readFromEnvProd(&env)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("APP_ENV is not set")
	}

	err := viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env, nil
}
