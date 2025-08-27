package models

import "gorm.io/gorm"

// ReadingProgress tracks user's progress in a novel.
type ReadingProgress struct {
	gorm.Model
	UserID    uint `gorm:"not null;uniqueIndex:idx_user_novel"`
	NovelID   uint `gorm:"not null;uniqueIndex:idx_user_novel"`
	ChapterID uint `gorm:"not null"`
	Position  int  `gorm:"not null"`
	User      User
	Novel     Novel
	Chapter   Chapter
}
