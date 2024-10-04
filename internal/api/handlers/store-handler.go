package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/api/repository"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	StoreRepo *repository.StoreRepository
}

func NewStoreHandler(repo *repository.StoreRepository) *StoreHandler {
	return &StoreHandler{StoreRepo: repo}
}

// Get all stores
func (h *StoreHandler) GetStores(c echo.Context) error {
	stores, err := h.StoreRepo.GetStores()
	if err != nil {
		log.Printf("Error fetching stores: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch stores")
	}

	return c.JSON(http.StatusOK, stores)
}

// Create a new store (Admin only)
func (h *StoreHandler) CreateStore(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*models.User)

	store := &models.Store{
		Name:        input.Name,
		Description: input.Description,
		UserID:      user.ID,
	}

	if err := h.StoreRepo.CreateStore(store); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, store)
}

// Get store by ID
func (h *StoreHandler) GetStoreByID(c echo.Context) error {
	storeID := c.Param("id")
	id, err := uuid.Parse(storeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid store ID")
	}

	store, err := h.StoreRepo.GetStoreByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Store not found")
	}

	return c.JSON(http.StatusOK, store)
}

// Update store (Admin only)
func (h *StoreHandler) UpdateStore(c echo.Context) error {
	storeID := c.Param("id")
	id, err := uuid.Parse(storeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid store ID")
	}

	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed")
	}

	store, err := h.StoreRepo.GetStoreByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Store not found")
	}

	store.Name = input.Name
	store.Description = input.Description

	if err := h.StoreRepo.UpdateStore(store); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update store")
	}

	return c.JSON(http.StatusOK, store)
}

// Delete store (Admin only)
func (h *StoreHandler) DeleteStore(c echo.Context) error {
	storeID := c.Param("id")
	id, err := uuid.Parse(storeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid store ID")
	}

	if err := h.StoreRepo.DeleteStore(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete store")
	}

	return c.NoContent(http.StatusNoContent)
}
