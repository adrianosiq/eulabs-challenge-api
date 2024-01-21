package handlers

import (
	"net/http"
	"strconv"

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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode product data")
	}

	createdProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create product")
	}

	return c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) Show(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing product ID")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get product")
	}

	return c.JSON(http.StatusOK, product)
}
