package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "github.com/IsaiasAraujo06/bike-inventory-api/internal/database"
	"github.com/IsaiasAraujo06/bike-inventory-api/internal/models"
)

type ProductHandler struct {
	repo *db.ProductRepository
}

// NewProductHandler cria novo handler
func NewProductHandler(repo *db.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// ErrorResponse estrutura de erro
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SuccessResponse estrutura de sucesso
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

// respondJSON envia resposta JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// respondError envia erro JSON
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}

// Create manipula POST /products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	var req models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	defer r.Body.Close()

	product, err := h.repo.Create(&req)
	if err != nil {
		// Erros de validação
		if err == models.ErrInvalidProductName ||
			err == models.ErrInvalidProductCategory ||
			err == models.ErrInvalidProductPrice {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	respondJSON(w, http.StatusCreated, SuccessResponse{
		Data:    product,
		Message: "Product created successfully",
	})
}

// GetByID manipula GET /products/:id
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Only GET method is allowed")
		return
	}

	// Extrai ID da URL
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 {
		respondError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.repo.GetByID(id)
	if err != nil {
		if err == models.ErrProductNotFound {
			respondError(w, http.StatusNotFound, "Product not found")
			return
		}

		if err == models.ErrInvalidProductName ||
			err == models.ErrInvalidProductCategory {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondError(w, http.StatusInternalServerError, "Failed to get product")
		return
	}

	respondJSON(w, http.StatusOK, SuccessResponse{
		Data: product,
	})
}
