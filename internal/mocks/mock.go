package mocks

import (
	"time"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAllProducts() ([]*models.Product, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

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
