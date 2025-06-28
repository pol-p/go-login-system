package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func VarEnv() *Config {
	err := godotenv.Load(filepath.Join("config", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port: os.Getenv("PORT"),
	}
}
