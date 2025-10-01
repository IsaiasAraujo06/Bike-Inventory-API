package main

import (
	"time"
)

// Criação da estrutura do produto
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
