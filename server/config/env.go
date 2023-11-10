package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	DSN string
	SERVER_PORT string
}

func LoadEnv() EnvVars {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_DSN")
	serverPort := os.Getenv("SERVER_PORT")

	return EnvVars{
		DSN: dsn,
		SERVER_PORT: serverPort,
	}
}