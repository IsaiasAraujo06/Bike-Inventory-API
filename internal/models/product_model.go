package models

import (
	"errors"
	"time"
)

// Representação do produto no inventário
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Payload de criação de um novo produto
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

// Método de validação dos dados do payload de criação
func (r *CreateProductRequest) Validate() error {
	if r.Name == "" {
		return errors.New("product name is required")
	}
	if r.Category == "" {
		return errors.New("product category is required")
	}
	if r.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	return nil
}

// Possíveis erros relacionados a produtos
var (
	ErrProductNotFound        = errors.New("product not found")
	ErrInvalidProductName     = errors.New("invalid product name")
	ErrInvalidProductCategory = errors.New("invalid product category")
	ErrInvalidProductPrice    = errors.New("invalid product price")
	ErrInsufficientStock      = errors.New("insufficient stock")
)
