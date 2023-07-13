package config

import (
	"os"
	"regexp"

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

func LoadEnv(fileName string) {
	re := regexp.MustCompile(`^(.*` + "twitter" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/` + fileName)
	if err != nil {
		godotenv.Load()
	}
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("postgresql://postgres:dreamer@20.203.31.58:5432/twitter?sslmode=disable"),
		},
		JWT: jwt{
			// ! chanage it later
			Secret: os.Getenv("dreamer"),
			Issuer: os.Getenv("twitter-clone"),
		},
	}
}
