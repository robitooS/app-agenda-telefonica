package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_USER   string
	DB_PASS   string
	DB_HOST   string
	DB_PORT   string
	DB_NAME   string
	DB_SOURCE string
	LOG_PATH  string
	API_PORT  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao carregar o arquivo .env: %v. Usando variáveis de ambiente.", err)
		// Continua o carregamento usando variáveis de ambiente do sistema,
		// o erro é registrado para debug, mas não impede a inicialização.
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	logPath := os.Getenv("LOG_PATH")
	apiPort := os.Getenv("API_PORT")

	dbSource := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	if logPath == "" {
		logPath = "logs/exclusao.log"
	}
	if apiPort == "" {
		apiPort = "8080"
	}

	return &Config{
		DB_USER:   dbUser,
		DB_PASS:   dbPass,
		DB_HOST:   dbHost,
		DB_PORT:   dbPort,
		DB_NAME:   dbName,
		DB_SOURCE: dbSource,
		LOG_PATH:  logPath,
		API_PORT:  apiPort,
	}, nil
}
