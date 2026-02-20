package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
)

// RunMigrations executa as migrações do banco de dados.
func RunMigrations(databaseURL string) error {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao banco de dados para migração: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("não foi possível conectar ao banco de dados (ping falhou): %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("falha ao criar a instância do driver de migração: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("falha ao criar a instância de migração: %w", err)
	}

	fmt.Println("Aplicando migrações...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("falha ao aplicar as migrações: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("falha ao verificar a versão da migração: %w", err)
	}

	fmt.Printf("Migrações aplicadas com sucesso. Versão atual: %d, Dirty: %v\n", version, dirty)
	return nil
}