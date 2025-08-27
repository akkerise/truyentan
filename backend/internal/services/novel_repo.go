package services

import (
	"github.com/truyentan/backend/internal/models"
	"gorm.io/gorm"
)

// NovelRepository handles database operations for novels.
type NovelRepository struct {
	db *gorm.DB
}

// NewNovelRepository creates a new NovelRepository.
func NewNovelRepository(db *gorm.DB) *NovelRepository {
	return &NovelRepository{db: db}
}

// ListNovels returns novels filtered by query parameters with pagination.
func (r *NovelRepository) ListNovels(page, limit int, query, genre, status string) ([]models.Novel, error) {
	var novels []models.Novel
	db := r.db.Model(&models.Novel{}).Preload("Genres")
	if query != "" {
		db = db.Where("title ILIKE ?", "%"+query+"%")
	}
	if status != "" {
		db = db.Where("status = ?", status)
	}
	if genre != "" {
		db = db.Joins("JOIN novel_genres ng ON ng.novel_id = novels.id").Joins("JOIN genres g ON g.id = ng.genre_id").Where("g.name = ?", genre)
	}
	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}
	if err := db.Find(&novels).Error; err != nil {
		return nil, err
	}
	return novels, nil
}

// GetNovel retrieves a novel by ID with genres and chapter count.
func (r *NovelRepository) GetNovel(id uint) (*models.Novel, int, error) {
	var novel models.Novel
	if err := r.db.Preload("Genres").First(&novel, id).Error; err != nil {
		return nil, 0, err
	}
	var count int64
	if err := r.db.Model(&models.Chapter{}).Where("novel_id = ?", id).Count(&count).Error; err != nil {
		return &novel, 0, err
	}
	return &novel, int(count), nil
}

// ListChapters returns chapters of a novel with pagination ordered by number.
func (r *NovelRepository) ListChapters(novelID uint, page, limit int) ([]models.Chapter, error) {
	var chapters []models.Chapter
	db := r.db.Where("novel_id = ?", novelID).Order("number")
	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}
	if err := db.Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}
