package postgres

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvCloudURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_CLOUD_URL")
}

func EnvHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_HOST")
}

func EnvUser() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_USER")
}

func EnvPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_PASSWORD")
}

func EnvDbName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_DBNAME")
}

func EnvPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POSTGRES_PORT")
}