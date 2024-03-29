package interfaces

import "github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"

type ProductRespositoryInterface interface {
	GetAll() ([]*models.Product, error)
	Create(product *models.Product) (*models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
	Delete(id int) error
}
