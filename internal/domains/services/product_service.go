package services

import (
	"github.com/adrianosiqe/eulabs-challenge-api/internal/core/interfaces"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
)

type ProductService struct {
	productRepository interfaces.ProductRespositoryInterface
}

func NewProductService(productRepository interfaces.ProductRespositoryInterface) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.productRepository.GetAll()
}
