package models

import "gorm.io/gorm"

// Chapter represents a single chapter in a novel.
type Chapter struct {
	gorm.Model
	NovelID uint   `gorm:"not null;uniqueIndex:idx_novel_chapter"`
	Number  int    `gorm:"not null;uniqueIndex:idx_novel_chapter"`
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text"`
	Novel   Novel
}
