package database

import (
	"context"

	"github.com/marius004/url-shortener/models"
)

type ShortUrlService struct {
	db *Database
}

func (s *ShortUrlService) GetByShortUrl(ctx context.Context, shortUrl string) (*models.ShortUrl, error) {
	var url *models.ShortUrl
	result := s.db.Conn.Where("short_url = ?", shortUrl).First(&shortUrl)

	return url, result.Error
}

func (s *ShortUrlService) Create(ctx context.Context, shortUrl *models.ShortUrl) error {
	result := s.db.Conn.Create(&shortUrl)
	return result.Error
}

func (s *ShortUrlService) Delete(ctx context.Context, id uint64) error {
	result := s.db.Conn.Delete(&models.ShortUrl{}, id)
	return result.Error
}

func (s *ShortUrlService) Update(ctx context.Context, id uint64, shortUrl *models.ShortUrl) error {
	result := s.db.Conn.Model(&models.ShortUrl{}).Where("id = ?", id).Update("views", shortUrl.Views)
	return result.Error
}

func (s *ShortUrlService) GetByUserId(ctx context.Context, userId uint64) ([]*models.ShortUrl, error) {
	var urls []*models.ShortUrl
	result := s.db.Conn.Where("user_id = ?", userId).Find(&urls)

	return urls, result.Error
}

func NewShortUrlService(db *Database) *ShortUrlService {
	return &ShortUrlService{
		db: db,
	}
}
