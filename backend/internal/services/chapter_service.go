package services

import (
	"github.com/truyentan/backend/internal/models"
	"gorm.io/gorm"
)

// ChapterService provides methods for retrieving chapters with navigation info.
type ChapterService struct {
	db *gorm.DB
}

// NewChapterService creates a new ChapterService.
func NewChapterService(db *gorm.DB) *ChapterService {
	return &ChapterService{db: db}
}

// GetByID fetches a chapter by id and returns chapter along with previous and next chapter IDs.
func (s *ChapterService) GetByID(id uint) (*models.Chapter, *uint, *uint, error) {
	var chapter models.Chapter
	if err := s.db.First(&chapter, id).Error; err != nil {
		return nil, nil, nil, err
	}
	prevID, nextID, err := s.navigation(&chapter)
	if err != nil {
		return nil, nil, nil, err
	}
	return &chapter, prevID, nextID, nil
}

// GetNext returns the chapter following the given chapter ID.
func (s *ChapterService) GetNext(id uint) (*models.Chapter, *uint, *uint, error) {
	var current models.Chapter
	if err := s.db.First(&current, id).Error; err != nil {
		return nil, nil, nil, err
	}
	var next models.Chapter
	if err := s.db.Where("novel_id = ? AND number = ?", current.NovelID, current.Number+1).First(&next).Error; err != nil {
		return nil, nil, nil, err
	}
	prevID, nextID, err := s.navigation(&next)
	if err != nil {
		return nil, nil, nil, err
	}
	return &next, prevID, nextID, nil
}

// GetPrev returns the chapter preceding the given chapter ID.
func (s *ChapterService) GetPrev(id uint) (*models.Chapter, *uint, *uint, error) {
	var current models.Chapter
	if err := s.db.First(&current, id).Error; err != nil {
		return nil, nil, nil, err
	}
	var prev models.Chapter
	if err := s.db.Where("novel_id = ? AND number = ?", current.NovelID, current.Number-1).First(&prev).Error; err != nil {
		return nil, nil, nil, err
	}
	prevID, nextID, err := s.navigation(&prev)
	if err != nil {
		return nil, nil, nil, err
	}
	return &prev, prevID, nextID, nil
}

// navigation determines the previous and next chapter IDs for a given chapter.
func (s *ChapterService) navigation(ch *models.Chapter) (*uint, *uint, error) {
	var prev models.Chapter
	if err := s.db.Select("id").Where("novel_id = ? AND number = ?", ch.NovelID, ch.Number-1).First(&prev).Error; err == nil {
		prevID := prev.ID
		var next models.Chapter
		if err := s.db.Select("id").Where("novel_id = ? AND number = ?", ch.NovelID, ch.Number+1).First(&next).Error; err == nil {
			nextID := next.ID
			return &prevID, &nextID, nil
		}
		return &prevID, nil, nil
	} else if err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}
	var next models.Chapter
	if err := s.db.Select("id").Where("novel_id = ? AND number = ?", ch.NovelID, ch.Number+1).First(&next).Error; err == nil {
		nextID := next.ID
		return nil, &nextID, nil
	} else if err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}
	return nil, nil, nil
}
