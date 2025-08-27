package models

import "gorm.io/gorm"

// Novel represents a story with multiple chapters.
type Novel struct {
	gorm.Model
	Title       string            `gorm:"type:varchar(255);not null"`
	Author      string            `gorm:"type:varchar(100);not null"`
	Description string            `gorm:"type:text"`
	Status      string            `gorm:"type:varchar(20);not null"`
	Chapters    []Chapter         `gorm:"constraint:OnDelete:CASCADE"`
	Genres      []Genre           `gorm:"many2many:novel_genres;constraint:OnDelete:CASCADE"`
	Favorites   []Favorite        `gorm:"constraint:OnDelete:CASCADE"`
	Progresses  []ReadingProgress `gorm:"constraint:OnDelete:CASCADE"`
}
