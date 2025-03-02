package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config struct holds all configuration values
type Rds struct {
	DSN string
}

type Server struct {
	Host string
	Port string
}
type Config struct {
	Rds    Rds
	Server Server
}

// LoadConfig loads environment variables from a .env file and returns the config
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	var config Config

	// Read the environment variables into the struct
	config.Rds.DSN = os.Getenv("DB_DSN")
	config.Server.Host = os.Getenv("SERVER_HOST")
	config.Server.Port = os.Getenv("SERVER_PORT")

	return &config, nil
}
