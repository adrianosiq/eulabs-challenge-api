package interfaces

import "github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"

type ProductRespositoryInterface interface {
	GetAll() ([]*models.Product, error)
}
