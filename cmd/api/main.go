package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "data.go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Stats()
	log.Println("Database connected")

	// Execução do SQL de migração
	sqlBytes, err := ioutil.ReadFile("migrations/001_create_products.sql")
	if err != nil { // Tratamento de erro na leitura do arquivo
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil { // Tratamento de erro na execução do SQL
		log.Fatal(err)
	}

	log.Println("Migration executed successfully")

}
