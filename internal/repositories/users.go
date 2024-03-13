package repositories

import (
	"github.com/b3nhard/chitempl/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(_db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: _db,
	}
}

func (r UserRepository) GetUser(_id int) *models.User {
	return nil
}
