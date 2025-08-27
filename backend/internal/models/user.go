package models

import "gorm.io/gorm"

// User represents application user.
type User struct {
	gorm.Model
	Name            string            `gorm:"type:varchar(100);not null"`
	Email           string            `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash    string            `gorm:"type:varchar(255);not null"`
	Favorites       []Favorite        `gorm:"constraint:OnDelete:CASCADE"`
	ReadingProgress []ReadingProgress `gorm:"constraint:OnDelete:CASCADE"`
}
