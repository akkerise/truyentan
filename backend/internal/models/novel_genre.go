package models

// NovelGenre is the join table between Novel and Genre.
type NovelGenre struct {
	NovelID uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}
