package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAllProducts() ([]*models.Product, error) {
	args := m.Called()
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*models.Product), args.Error(1)
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

func TestIndex(t *testing.T) {
	t.Run("should returns 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products")

		mockProductService := &MockProductService{}
		mockProductService.On("GetAllProducts").Return(MockProducts, nil).Once()
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Index(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var products []models.Product
			json.Unmarshal(rec.Body.Bytes(), &products)

			assert.Equal(t, 2, len(products))
			assert.Equal(t, uint(1), products[0].ID)
			assert.Equal(t, "Bulbasaur", products[0].Title)
			assert.Contains(t, products[0].Description, "There is a plant seed")
			assert.Equal(t, 99.99, products[0].Price)
			assert.Equal(t, uint(2), products[1].ID)
			assert.Equal(t, "Charmander", products[1].Title)
			assert.Contains(t, products[1].Description, "It has a preference for hot")
			assert.Equal(t, 1093.45, products[1].Price)
		}
	})
}

func TestCreate(t *testing.T) {
	var productJSON = `{"title":"Charmander","description":"It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.", "price": 1093.45}`

	t.Run("should returns 201", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products")

		var productBind models.Product
		json.Unmarshal([]byte(productJSON), &productBind)
		mockProductService := &MockProductService{}
		mockProductService.On("CreateProduct", &productBind).Return(MockProducts[1], nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var product models.Product
			json.Unmarshal(rec.Body.Bytes(), &product)

			assert.Equal(t, uint(2), product.ID)
			assert.Equal(t, "Charmander", product.Title)
			assert.Contains(t, product.Description, "It has a preference")
			assert.Equal(t, 1093.45, product.Price)
		}
	})
}
