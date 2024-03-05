package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT        string
	LOG_LEVEL       string
	LOG_OUTPUT      string
	ALLOWED_ORIGINS string

	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_DB_NAME  string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	API_PORT = os.Getenv("API_PORT")
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	LOG_OUTPUT = os.Getenv("LOG_OUTPUT")
	ALLOWED_ORIGINS = os.Getenv("ALLOWED_ORIGINS")

	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
}
