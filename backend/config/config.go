package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	ServerPort  int
}


// on 
func LoadConfig(env string) (*Config, error) {
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("Error loading .env.%s file", env)
	}
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Invalid server port: %v", err)
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  serverPort,
	}, nil
}
