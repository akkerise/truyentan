package models

import "gorm.io/gorm"

// Genre categorizes novels.
type Genre struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Novels []Novel `gorm:"many2many:novel_genres"`
}
