package repository

import (
	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/jmoiron/sqlx"
)

type StoreRepository struct {
	DB *sqlx.DB
}

func NewStoreRepository(db *sqlx.DB) *StoreRepository {
	return &StoreRepository{DB: db}
}

func (r *StoreRepository) GetStores() ([]models.Store, error) {
	stores := []models.Store{}
	query := `SELECT * FROM stores`
	err := r.DB.Select(&stores, query)
	return stores, err
}

func (r *StoreRepository) CreateStore(store *models.Store) error {
	query := `
        INSERT INTO stores (name, description, user_id)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at`
	err := r.DB.QueryRowx(query, store.Name, store.Description, store.UserID).
		Scan(&store.ID, &store.CreatedAt, &store.UpdatedAt)
	return err
}

func (r *StoreRepository) GetStoreByID(id uuid.UUID) (*models.Store, error) {
	var store models.Store
	query := `SELECT * FROM stores WHERE id = $1`
	err := r.DB.Get(&store, query, id)
	return &store, err
}

func (r *StoreRepository) UpdateStore(store *models.Store) error {
	query := `
        UPDATE stores
        SET name = $1, description = $2, updated_at = NOW()
        WHERE id = $3
        RETURNING updated_at`
	err := r.DB.QueryRowx(query, store.Name, store.Description, store.ID).
		Scan(&store.UpdatedAt)
	return err
}

func (r *StoreRepository) DeleteStore(id uuid.UUID) error {
	query := `DELETE FROM stores WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
