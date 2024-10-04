package repository

import (
	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetCategories() ([]models.Category, error) {
	categories := []models.Category{}
	query := `SELECT * FROM categories`
	err := r.DB.Select(&categories, query)
	return categories, err
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	query := `
        INSERT INTO categories (name, description, store_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`
	err := r.DB.QueryRowx(query, category.Name, category.Description, category.ParentID, category.StoreID).
		Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
	return err
}

func (r *CategoryRepository) GetCategoryByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	query := `SELECT * FROM categories WHERE id = $1`
	err := r.DB.Get(&category, query, id)
	return &category, err
}

func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	query := `
        UPDATE categories
        SET name = $1, description = $2, updated_at = NOW()
        WHERE id = $3
        RETURNING updated_at`

	return r.DB.QueryRow(query, category.Name, category.Description, category.ID).
		Scan(&category.UpdatedAt)
}

func (r *CategoryRepository) DeleteCategory(id uuid.UUID) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
