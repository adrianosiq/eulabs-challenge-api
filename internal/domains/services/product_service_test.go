package services

import (
	"fmt"
	"testing"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	t.Run("should return a list the products", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("GetAll").Return(mocks.MockProducts, nil)

		productService := NewProductService(mockProductRepository)
		products, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.True(t, len(products) == 2)
		assert.Equal(t, mocks.MockProducts[0].ID, products[0].ID)
		assert.Equal(t, mocks.MockProducts[0].Title, products[0].Title)
		assert.Equal(t, mocks.MockProducts[0].Description, products[0].Description)
		assert.Equal(t, mocks.MockProducts[0].Price, products[0].Price)
		assert.Equal(t, mocks.MockProducts[1].ID, products[1].ID)
		assert.Equal(t, mocks.MockProducts[1].Title, products[1].Title)
		assert.Equal(t, mocks.MockProducts[1].Description, products[1].Description)
		assert.Equal(t, mocks.MockProducts[1].Price, products[1].Price)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		var mockEmptyProducts []*models.Product
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("GetAll").Return(mockEmptyProducts, nil)

		productService := NewProductService(mockProductRepository)
		products, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.Empty(t, products)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("GetAll").Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.GetAllProducts()

		assert.Error(t, err)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestCreateProducts(t *testing.T) {
	var mockCreateProduct = models.Product{Title: mocks.MockProducts[0].Title, Description: mocks.MockProducts[0].Description, Price: mocks.MockProducts[0].Price}

	t.Run("should return an product", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Create", &mockCreateProduct).Return(mocks.MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.CreateProduct(&mockCreateProduct)

		assert.NoError(t, err)
		assert.Equal(t, mocks.MockProducts[0].ID, product.ID)
		assert.Equal(t, mocks.MockProducts[0].Title, product.Title)
		assert.Equal(t, mocks.MockProducts[0].Description, product.Description)
		assert.Equal(t, mocks.MockProducts[0].Price, product.Price)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Create", &mockCreateProduct).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.CreateProduct(&mockCreateProduct)

		assert.Error(t, err)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestGetProductByID(t *testing.T) {
	t.Run("should return the product", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("GetByID", 1).Return(mocks.MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.GetProductByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mocks.MockProducts[0].ID, product.ID)
		assert.Equal(t, mocks.MockProducts[0].Title, product.Title)
		assert.Equal(t, mocks.MockProducts[0].Description, product.Description)
		assert.Equal(t, mocks.MockProducts[0].Price, product.Price)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("GetByID", 1).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.GetProductByID(1)

		assert.Error(t, err)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("should return the product", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Update", mocks.MockProducts[0]).Return(mocks.MockProducts[0], nil)

		productService := NewProductService(mockProductRepository)
		product, err := productService.UpdateProduct(mocks.MockProducts[0])

		assert.NoError(t, err)
		assert.Equal(t, mocks.MockProducts[0].ID, product.ID)
		assert.Equal(t, mocks.MockProducts[0].Title, product.Title)
		assert.Equal(t, mocks.MockProducts[0].Description, product.Description)
		assert.Equal(t, mocks.MockProducts[0].Price, product.Price)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Update", mocks.MockProducts[0]).Return(nil, fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		_, err := productService.UpdateProduct(mocks.MockProducts[0])

		assert.Error(t, err)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Delete", 1).Return(nil)

		productService := NewProductService(mockProductRepository)
		err := productService.DeleteProduct(1)

		assert.NoError(t, err)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should return an error", func(t *testing.T) {
		mockProductRepository := &mocks.MockProductRepository{}
		mockProductRepository.On("Delete", 1).Return(fmt.Errorf("some error"))

		productService := NewProductService(mockProductRepository)
		err := productService.DeleteProduct(1)

		assert.Error(t, err)
		mockProductRepository.AssertExpectations(t)
	})
}
