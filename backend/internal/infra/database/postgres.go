package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Driver do postgres
)

func NewConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir o driver de banco: %v", err)
	}

	db.SetMaxOpenConns(10)                 
	db.SetMaxIdleConns(5)                  
	db.SetConnMaxLifetime(time.Minute * 5) 

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("não foi possível efetuar uma conexão com o banco: %v", err)
	}

	return db, nil
}