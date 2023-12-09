package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	DSN string
	SERVER_PORT string
	GOOGLE_KEY string
	GOOGLE_CLIENT_ID string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_CALLBACK_URL string
}

func LoadEnv() EnvVars {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_DSN")
	serverPort := os.Getenv("SERVER_PORT")

	googleKey := os.Getenv("GOOGLE_KEY")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleCallbackUrl := os.Getenv("GOOGLE_CALLBACK_URL")

	return EnvVars{
		DSN: dsn,
		SERVER_PORT: serverPort,
		GOOGLE_KEY: googleKey,
		GOOGLE_CLIENT_ID: googleClientId,
		GOOGLE_CLIENT_SECRET: googleClientSecret,
		GOOGLE_CALLBACK_URL: googleCallbackUrl,
	}
}