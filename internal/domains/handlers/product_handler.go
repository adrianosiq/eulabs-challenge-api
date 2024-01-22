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

func (h *ProductHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing product ID")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}

	var updateProduct models.Product
	err = c.Bind(&updateProduct)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode product data")
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get product")
	}

	product.Title = updateProduct.Title
	product.Description = updateProduct.Description
	product.Price = updateProduct.Price

	updatedProduct, err := h.productService.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update product")
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing product ID")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}

	err = h.productService.DeleteProduct(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete product")
	}

	return c.JSON(http.StatusNoContent, nil)
}
