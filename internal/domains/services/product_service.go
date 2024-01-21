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

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepository.Create(product)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}
