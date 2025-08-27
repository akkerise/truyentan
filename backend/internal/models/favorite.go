package models

import "gorm.io/gorm"

// Favorite links a user with their favorite novels.
type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"not null;uniqueIndex:idx_user_novel"`
	NovelID uint `gorm:"not null;uniqueIndex:idx_user_novel"`
	User    User
	Novel   Novel
}
