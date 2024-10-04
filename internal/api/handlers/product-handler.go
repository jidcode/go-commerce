package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/api/repository"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductRepo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{ProductRepo: repo}
}

// Get all products
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.ProductRepo.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch products")
	}

	return c.JSON(http.StatusOK, products)
}

// Create a new product (Admin only)
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var input struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required"`
		Quantity    int     `json:"quantity"`
		CategoryID  string  `json:"category_id" validate:"required"`
		StoreID     string  `json:"store_id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		log.Printf("Error binding product: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed")
	}

	categoryID, err := uuid.Parse(input.CategoryID)
	storeID, err := uuid.Parse(input.StoreID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		CategoryID:  categoryID,
		StoreID:     storeID,
	}

	if err := h.ProductRepo.CreateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create product")
	}

	return c.JSON(http.StatusCreated, product)
}

// Get product by ID
func (h *ProductHandler) GetProductByID(c echo.Context) error {
	productID := c.Param("id")
	id, err := uuid.Parse(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	product, err := h.ProductRepo.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

// Update product (Admin only)
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productID := c.Param("id")
	id, err := uuid.Parse(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	var input struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required"`
		Quantity    int     `json:"quantity"`
		CategoryID  string  `json:"category_id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	product, err := h.ProductRepo.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	categoryID, err := uuid.Parse(input.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Quantity = input.Quantity
	product.CategoryID = categoryID

	if err := h.ProductRepo.UpdateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update product")
	}

	return c.JSON(http.StatusOK, product)
}

// Delete product (Admin only)
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productID := c.Param("id")
	id, err := uuid.Parse(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	if err := h.ProductRepo.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete product")
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}
