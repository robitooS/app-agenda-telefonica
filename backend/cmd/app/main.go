package main

import ("database/sql"; "fmt"; "log"; "github.com/gin-gonic/gin"; _ "github.com/lib/pq"; "github.com/golang-migrate/migrate/v4"; _ "github.com/golang-migrate/migrate/v4/database/postgres"; _ "github.com/golang-migrate/migrate/v4/source/file"; "github.com/robitooS/backend/internal/config"; "github.com/robitooS/backend/internal/handler"; "github.com/robitooS/backend/internal/repository"; "github.com/robitooS/backend/internal/service")

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Falha ao carregar as configurações: %v", err)
	}

	// Conecta ao banco de dados
	db, err := sql.Open("postgres", cfg.DB_SOURCE)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Verifica a conexão com o banco de dados
	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}
	log.Println("Conectado ao banco de dados com sucesso!")

	// Executa as migrações
	m, err := migrate.New(
		"file://migrations",
		cfg.DB_SOURCE)
	if err != nil {
		log.Fatalf("Erro ao criar a instância de migração: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao executar as migrações: %v", err)
	}
	log.Println("Migrações aplicadas com sucesso!")

	// Inicializa o repositório, serviço e handler
	contatoRepo := repository.NewContatoPostgres(db)
	contatoService := service.NewContatoService(contatoRepo, cfg.LOG_PATH)
	contatoHandler := handler.NewContatoHandler(contatoService)

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
