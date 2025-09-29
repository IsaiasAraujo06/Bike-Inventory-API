package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Busca a porta do ambiente ou usa 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Handler simples pra testar
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Bike Inventory API is running", "status": "ok"}`)
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "healthy"}`)
	})

	addr := ":" + port
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📍 http://localhost:%s", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
