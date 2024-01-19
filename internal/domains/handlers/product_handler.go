package handlers

import (
	"net/http"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/core/interfaces"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/labstack/echo"
)

type ProductHandler struct {
	productService interfaces.ProductServiceInterface
}

type ProductInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func NewProductHandler(productService interfaces.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Index(c echo.Context) error {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to list the products")
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Create(c echo.Context) error {
	var product models.Product

	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode user data")
	}

	createdProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, createdProduct)
}
