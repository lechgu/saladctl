package config

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/samber/do"
)

type Config struct {
	BaseURL  string `env:"BASE_URL,required"`
	Email    string `env:"EMAIL,required"`
	Password string `env:"PASSWORD,required"`
}

func New(di *do.Injector) (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load()
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
