package config

import (
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	URL string
}

type Config struct {
	Database database
	JWT      jwt
}

type jwt struct {
	Secret string
	Issuer string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
		JWT: jwt{
			Secret: os.Getenv("JWT_SECRET"),
			Issuer: os.Getenv("DOMAIN"),
		},
	}
}
