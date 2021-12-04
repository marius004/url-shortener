package services

import (
	"context"

	"github.com/marius004/url-shortener/models"
)

type UserService interface {
	GetById(ctx context.Context, id uint64) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	Create(ctx context.Context, user *models.User) error
}
