package interfaces

import "github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"

type ProductServiceInterface interface {
	GetAllProducts() ([]*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
}
