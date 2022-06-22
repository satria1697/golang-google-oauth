package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectUri  string `env:"REDIRECT_URI"`
	Port         string `env:"PORT"`
}

func InitConfig() *Config {
	godotenv.Read()
	return &Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectUri:  os.Getenv("REDIRECT_URI"),
		Port:         os.Getenv("PORT"),
	}
}
