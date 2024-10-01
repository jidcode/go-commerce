package repository

import (
	"github.com/google/uuid"
	"github.com/jidcode/go-commerce/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password, role)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, created_at, updated_at`

	err := repo.DB.QueryRowx(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE email = $1`
	err := repo.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (repo *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE id = $1`
	err := repo.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
