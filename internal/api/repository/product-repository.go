package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	products := []models.Product{}
	query := `SELECT * FROM products`
	err := r.DB.Select(&products, query)
	return products, err
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	product.ID = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	query := `INSERT INTO products (id, name, description, price, quantity, category_id, store_id, created_at, updated_at)
              VALUES (:id, :name, :description, :price, :quantity, :category_id, :store_id, :created_at, :updated_at)`

	_, err := r.DB.NamedExec(query, product)
	return err
}

func (r *ProductRepository) GetProductByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.DB.Get(&product, query, id)
	return &product, err
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	product.UpdatedAt = time.Now()

	query := `UPDATE products 
			  SET name = :name, description = :description, price = :price, quantity = :quantity, 
			  category_id = :category_id, updated_at = :updated_at 
			  WHERE id = :id`
	_, err := r.DB.NamedExec(query, product)
	return err
}

func (r *ProductRepository) DeleteProduct(id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
