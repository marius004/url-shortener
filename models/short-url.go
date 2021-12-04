package models

import "gorm.io/gorm"

type ShortUrl struct {
	gorm.Model
	UserID uint

	ShortUrl string `gorm:"unique;not null"`
	LongURL  string `gorm:"unique;not null"`

	Views int `gorm:"default:0"`
}
