package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll() ([]*models.Product, error) {
	args := m.Called()
	return args.Get(0).([]*models.Product), args.Error(1)
}

var MockProducts = []*models.Product{
	{
		ID:          1,
		Title:       "Bulbasaur",
		Description: "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.",
		Price:       99.99,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "Charmander",
		Description: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",
		Price:       1093.45,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func TestGetAllProducts(t *testing.T) {
	t.Run("should return a list the products", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetAll").Return(MockProducts, nil)

		productService := NewProductService(mockProductRepository)
		products, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.True(t, len(products) == 2)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		var mockEmptyProducts []*models.Product
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetAll").Return(mockEmptyProducts, nil)

		productService := NewProductService(mockProductRepository)
		products, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.Empty(t, products)
	})

	t.Run("should return an error", func(t *testing.T) {
		var mockEmptyProducts []*models.Product
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetAll").Return(mockEmptyProducts, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.GetAllProducts()

		assert.Error(t, err)
	})
}
