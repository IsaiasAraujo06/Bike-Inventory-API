package database

import (
	"database/sql"
	"fmt"

	"github.com/IsaiasAraujo06/bike-inventory-api/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository cria novo repository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create insere novo produto
func (r *ProductRepository) Create(req *models.CreateProductRequest) (*models.Product, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	query := `
		INSERT INTO products (name, category, price, stock, description)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, req.Name, req.Category, req.Price, req.Stock, req.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return r.GetByID(int(id))
}

// GetByID busca produto por ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT id, name, category, price, stock, description, created_at, updated_at
		FROM products
		WHERE id = ?
	`

	var product models.Product
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Category,
		&product.Price,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, models.ErrProductNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &product, nil
}
