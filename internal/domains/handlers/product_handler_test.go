package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/mocks"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}

func TestIndex(t *testing.T) {
	t.Run("should returns 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetAllProducts").Return(mocks.MockProducts, nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Index(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var products []models.Product
			json.Unmarshal(rec.Body.Bytes(), &products)

			assert.Equal(t, 2, len(products))
			assert.Equal(t, mocks.MockProducts[0].ID, products[0].ID)
			assert.Equal(t, mocks.MockProducts[0].Title, products[0].Title)
			assert.Equal(t, mocks.MockProducts[0].Description, products[0].Description)
			assert.Equal(t, mocks.MockProducts[0].Price, products[0].Price)
			assert.Equal(t, mocks.MockProducts[1].ID, products[1].ID)
			assert.Equal(t, mocks.MockProducts[1].Title, products[1].Title)
			assert.Equal(t, mocks.MockProducts[1].Description, products[1].Description)
			assert.Equal(t, mocks.MockProducts[1].Price, products[1].Price)

			mockProductService.AssertExpectations(t)
		}
	})

	t.Run("should returns 500", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetAllProducts").Return(nil, fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Index(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=500, message=Failed to list the products")
		mockProductService.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var productJSON = `{"title":"Charmander","description":"It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.", "price": 1093.45}`

	t.Run("should returns 201", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		var productBind models.Product
		json.Unmarshal([]byte(productJSON), &productBind)
		mockProductService := &mocks.MockProductService{}
		mockProductService.On("CreateProduct", &productBind).Return(mocks.MockProducts[1], nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var product models.Product
			json.Unmarshal(rec.Body.Bytes(), &product)

			assert.Equal(t, mocks.MockProducts[1].ID, product.ID)
			assert.Equal(t, mocks.MockProducts[1].Title, product.Title)
			assert.Equal(t, mocks.MockProducts[1].Description, product.Description)
			assert.Equal(t, mocks.MockProducts[1].Price, product.Price)

			mockProductService.AssertExpectations(t)
		}
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)
		err := productHandler.Create(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Failed to decode product data")
	})

	t.Run("should returns 422", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", strings.NewReader(`{"title":"Charmander"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)
		err := productHandler.Create(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=422, message=Key: 'Product.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'Product.Price' Error:Field validation for 'Price' failed on the 'required' tag")
	})

	t.Run("should returns 500", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		var productBind models.Product
		json.Unmarshal([]byte(productJSON), &productBind)
		mockProductService := &mocks.MockProductService{}
		mockProductService.On("CreateProduct", &productBind).Return(nil, fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Create(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=500, message=Failed to create product")
		mockProductService.AssertExpectations(t)
	})
}

func TestShow(t *testing.T) {
	t.Run("should returns 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetProductByID", 1).Return(mocks.MockProducts[0], nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Show(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var product models.Product
			json.Unmarshal(rec.Body.Bytes(), &product)

			assert.Equal(t, mocks.MockProducts[0].ID, product.ID)
			assert.Equal(t, mocks.MockProducts[0].Title, product.Title)
			assert.Equal(t, mocks.MockProducts[0].Description, product.Description)
			assert.Equal(t, mocks.MockProducts[0].Price, product.Price)

			mockProductService.AssertExpectations(t)
		}
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Show(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Missing product ID")
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Show(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Invalid product ID")
	})

	t.Run("should returns 404", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetProductByID", 1).Return(nil, fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Show(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=404, message=Failed to get product")
		mockProductService.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	var productJSON = `{"title":"Charmander","description":"It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.", "price": 1093.45}`
	var updatedProduct = models.Product{
		ID:          mocks.MockProducts[0].ID,
		Title:       "Charmander",
		Description: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",
		Price:       1093.45,
		CreatedAt:   mocks.MockProducts[0].CreatedAt,
		UpdatedAt:   mocks.MockProducts[0].UpdatedAt,
		DeletedAt:   mocks.MockProducts[0].DeletedAt,
	}

	t.Run("should returns 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/:id", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetProductByID", 1).Return(mocks.MockProducts[0], nil)
		mockProductService.On("UpdateProduct", &updatedProduct).Return(&updatedProduct, nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Update(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var product models.Product
			json.Unmarshal(rec.Body.Bytes(), &product)

			assert.Equal(t, updatedProduct.ID, product.ID)
			assert.Equal(t, updatedProduct.Title, product.Title)
			assert.Equal(t, updatedProduct.Description, product.Description)
			assert.Equal(t, updatedProduct.Price, product.Price)

			mockProductService.AssertExpectations(t)
		}
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/:id", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Update(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Missing product ID")
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/:id", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Update(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Invalid product ID")
	})

	t.Run("should returns 404", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/:id", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetProductByID", 1).Return(nil, fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Update(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=404, message=Failed to get product")
		mockProductService.AssertExpectations(t)
	})

	t.Run("should returns 500", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products/:id", strings.NewReader(productJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		var productBind models.Product
		json.Unmarshal([]byte(productJSON), &productBind)
		mockProductService := &mocks.MockProductService{}
		mockProductService.On("GetProductByID", 1).Return(mocks.MockProducts[0], nil)
		mockProductService.On("UpdateProduct", &updatedProduct).Return(nil, fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Update(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=500, message=Failed to update product")
		mockProductService.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should returns 204", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("DeleteProduct", 1).Return(nil)
		productHandler := NewProductHandler(mockProductService)

		if assert.NoError(t, productHandler.Delete(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
			mockProductService.AssertExpectations(t)
		}
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Delete(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Missing product ID")
	})

	t.Run("should returns 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("invalid_id")

		mockProductService := &mocks.MockProductService{}
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Delete(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=400, message=Invalid product ID")
	})

	t.Run("should returns 500", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		mockProductService := &mocks.MockProductService{}
		mockProductService.On("DeleteProduct", 1).Return(fmt.Errorf("some error"))
		productHandler := NewProductHandler(mockProductService)

		err := productHandler.Delete(c)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "code=500, message=Failed to delete product")
		mockProductService.AssertExpectations(t)
	})
}
