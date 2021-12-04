package database

import (
	"context"

	"github.com/marius004/url-shortener/models"
)

// implements internal.UserService
type UserService struct {
	db *Database
}

func (s *UserService) GetById(ctx context.Context, id uint64) (*models.User, error) {
	var user *models.User
	result := s.db.Conn.First(&user, id)

	return user, result.Error
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	result := s.db.Conn.Where("username = ?", username).First(&user)

	return user, result.Error
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	result := s.db.Conn.Create(&user)
	return result.Error
}

func NewUserService(db *Database) *UserService {
	return &UserService{
		db: db,
	}
}
