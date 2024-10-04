package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/api/repository"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryHandler(repo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{CategoryRepo: repo}
}

// Get all categories
func (h *CategoryHandler) GetCategories(c echo.Context) error {
	categories, err := h.CategoryRepo.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch categories")
	}

	return c.JSON(http.StatusOK, categories)
}

// Create a new category (Admin only)
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var input struct {
		Name        string    `json:"name" validate:"required"`
		Description string    `json:"description"`
		StoreID     uuid.UUID `json:"store_id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed")
	}

	category := &models.Category{
		Name:        input.Name,
		Description: input.Description,
		StoreID:     input.StoreID,
	}

	if err := h.CategoryRepo.CreateCategory(category); err != nil {
		log.Printf("Error creating category: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to create category")
	}

	return c.JSON(http.StatusCreated, category)
}

// Get category by ID
func (h *CategoryHandler) GetCategoryByID(c echo.Context) error {
	categoryID := c.Param("id")
	id, err := uuid.Parse(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	category, err := h.CategoryRepo.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	return c.JSON(http.StatusOK, category)
}

// Update category (Admin only)
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	categoryID := c.Param("id")
	id, err := uuid.Parse(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	category, err := h.CategoryRepo.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	category.Name = input.Name
	category.Description = input.Description

	if err := h.CategoryRepo.UpdateCategory(category); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update category")
	}

	return c.JSON(http.StatusOK, category)
}

// Delete category (Admin only)
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	categoryID := c.Param("id")
	id, err := uuid.Parse(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	if err := h.CategoryRepo.DeleteCategory(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete category")
	}

	return c.NoContent(http.StatusNoContent)
}
