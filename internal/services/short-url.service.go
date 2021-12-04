package services

import (
	"context"

	"github.com/marius004/url-shortener/models"
)

type ShortUrlService interface {
	GetByShortUrl(ctx context.Context, shortUrl string) (*models.ShortUrl, error)
	GetByUserId(ctx context.Context, userId uint64) ([]*models.ShortUrl, error)

	Create(ctx context.Context, shortUrl *models.ShortUrl) error

	Delete(ctx context.Context, id uint64) error

	Update(ctx context.Context, id uint64, shortUrl *models.ShortUrl) error
}
