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
	API_PORT  string
	DelLogPath string // Caminho para o log de deleções (arquivo .txt)
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
	apiPort := os.Getenv("API_PORT")
	delLogPath := os.Getenv("DEL_LOG_PATH")

	dbSource := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	if apiPort == "" {
		apiPort = "8080"
	}
	if delLogPath == "" {
		delLogPath = "logs/deleted_contacts.txt" 
	}

	return &Config{
		DB_USER:   dbUser,
		DB_PASS:   dbPass,
		DB_HOST:   dbHost,
		DB_PORT:   dbPort,
		DB_NAME:   dbName,
		DB_SOURCE: dbSource,
		API_PORT:  apiPort,
		DelLogPath: delLogPath,
	}, nil
}
