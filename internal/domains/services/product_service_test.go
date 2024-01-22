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
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetByID(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
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
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetAll").Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.GetAllProducts()

		assert.Error(t, err)
	})
}

func TestCreateProducts(t *testing.T) {
	var mockCreateProduct = models.Product{Title: MockProducts[0].Title, Description: MockProducts[0].Description, Price: MockProducts[0].Price}

	t.Run("should return an product", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Create", &mockCreateProduct).Return(MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.CreateProduct(&mockCreateProduct)

		assert.NoError(t, err)
		assert.Equal(t, MockProducts[0].ID, product.ID)
		assert.Equal(t, MockProducts[0].Title, product.Title)
		assert.Equal(t, MockProducts[0].Description, product.Description)
		assert.Equal(t, MockProducts[0].Price, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Create", &mockCreateProduct).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.CreateProduct(&mockCreateProduct)

		assert.Error(t, err)
	})
}

func TestGetProductByID(t *testing.T) {
	t.Run("should return the product", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetByID", 1).Return(MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.GetProductByID(1)

		assert.NoError(t, err)
		assert.Equal(t, MockProducts[0].ID, product.ID)
		assert.Equal(t, MockProducts[0].Title, product.Title)
		assert.Equal(t, MockProducts[0].Description, product.Description)
		assert.Equal(t, MockProducts[0].Price, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("GetByID", 1).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.GetProductByID(1)

		assert.Error(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("should return the product", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Update", MockProducts[0]).Return(MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.UpdateProduct(MockProducts[0])

		assert.NoError(t, err)
		assert.Equal(t, MockProducts[0].ID, product.ID)
		assert.Equal(t, MockProducts[0].Title, product.Title)
		assert.Equal(t, MockProducts[0].Description, product.Description)
		assert.Equal(t, MockProducts[0].Price, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Update", MockProducts[0]).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.UpdateProduct(MockProducts[0])

		assert.Error(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Delete", 1).Return(nil)

		productService := NewProductService(mockProductRepository)
		err := productService.DeleteProduct(1)

		assert.NoError(t, err)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &MockProductRepository{}
		mockProductRepository.On("Delete", 1).Return(fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		err := productService.DeleteProduct(1)

		assert.Error(t, err)
	})
}
