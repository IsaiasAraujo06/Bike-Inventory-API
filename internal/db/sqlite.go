package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Função de criação do banco de dados SQLite
func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "bike-inventory-api/internal/db/sqlite.go")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
