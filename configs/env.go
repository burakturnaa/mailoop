package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvServerPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	serverPort := os.Getenv("SERVER_PORT")
	return serverPort
}

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoIRU := os.Getenv("MONGOURI")
	return mongoIRU
}

func EnvSMTP() map[string]string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	var smtp = map[string]string{
		"host":     os.Getenv("SMTP_HOST"),
		"port":     os.Getenv("SMTP_PORT"),
		"username": os.Getenv("SMTP_USERNAME"),
		"email":    os.Getenv("SMTP_EMAIL"),
		"password": os.Getenv("SMTP_PASSWORD"),
	}
	return smtp
}
