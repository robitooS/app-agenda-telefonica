package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robitooS/backend/internal/config"
	"github.com/robitooS/backend/internal/handler"
	"github.com/robitooS/backend/internal/infra/database"
	"github.com/robitooS/backend/internal/repository"
	"github.com/robitooS/backend/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Falha ao carregar as configurações: %v", err)
	}

	// Conecta ao banco de dados
	db, err := database.NewConnection(cfg.DB_SOURCE)
	if err != nil {
		log.Fatalf("Erro ao criar conexão com o banco de dados: %v", err);
	}
	defer db.Close();

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Erro ao executar as migrations: %v", err);
	}

	// Inicializa o repositório, serviço e handler
	contatoRepo := repository.NewContatoPostgres(db)
	contatoService := service.NewContatoService(contatoRepo)
	contatoHandler := handler.NewContatoHandler(contatoService, cfg.DelLogPath)

	// Configura o roteador Gin
	router := gin.Default()

	// Middleware CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	contatoHandler.RegisterRoutes(router)

	// Inicia o servidor
	log.Printf("Servidor iniciando na porta %s", cfg.API_PORT)
	if err := router.Run(fmt.Sprintf(":%s", cfg.API_PORT)); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
