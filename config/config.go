package config

import (
	"os"

	"github.com/joho/godotenv"
)

// type database struct {
// 	URL string
// }

// type Config struct {
// 	Database database
// }

// func New() *Config {
// 	godotenv.Load()
// 	return &Config{
// 		Database: database{
// 			URL: "postgres://postgres:@127.0.0.1:5432/twitter?sslmode=disable",
// 			// os.Getenv("DATABASE_URL"),
// 		},
// 	}
// }

type database struct {
	URL string
}

type Config struct {
	Database database
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}
}
