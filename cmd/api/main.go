package main

import (
	"log"
	"net/http"

	"github.com/IsaiasAraujo06/bike-inventory-api/internal/database"
	"github.com/IsaiasAraujo06/bike-inventory-api/internal/handlers"
)

func main() {
	// Inicializa database
	db, err := database.New("data.db")
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer db.Close()
	log.Println("✅ Database connected")

	// Executa migrations
	if err := db.RunMigrations(); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}
	log.Println("✅ Migrations executed successfully")

	// Inicializa repository e handler
	productRepo := database.NewProductRepository(db.DB)
	productHandler := handlers.NewProductHandler(productRepo)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "database": "connected"}`))
	})

	// Rotas de produtos
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			productHandler.Create(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", productHandler.GetByID)

	// Inicia servidor
	port := "8080"
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📍 http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
