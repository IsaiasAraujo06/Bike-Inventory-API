package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Estrutura que encapsula a conexão com o banco de dados SQLite
type DB struct {
	*sql.DB
}

// Método de inicialização da conexão com o banco de dados SQLite
func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite", dbPath) // Declaração do driver SQLite

	// Verificação de erros na abertura da conexão
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{db}, nil // Retorna a instância do DB
}

// Método de execução de migrations
func (db *DB) RunMigrations() error {
	sqlBytes, err := os.ReadFile("migrations/001_create_products.sql")

	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}

// Método de fechamento da conexão com o banco de dados
func (db *DB) Close() error {
	return db.DB.Close()
}
