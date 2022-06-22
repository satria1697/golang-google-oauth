package utils

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectUri  string `env:"REDIRECT_URI"`
}

func InitConfig() Config {
	var cfg Config
	cleanenv.ReadConfig(".env", &cfg)
	return cfg
}
