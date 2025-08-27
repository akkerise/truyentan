package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/truyentan/backend/internal/services"
)

// ProgressHandler handles reading progress endpoints.
type ProgressHandler struct {
	svc *services.ProgressService
}

// NewProgressHandler creates a new ProgressHandler.
func NewProgressHandler(svc *services.ProgressService) *ProgressHandler {
	return &ProgressHandler{svc: svc}
}

type progressSaveRequest struct {
	NovelID        uint `json:"novelId" binding:"required"`
	ChapterID      uint `json:"chapterId" binding:"required"`
	PositionOffset int  `json:"positionOffset" binding:"required"`
}

// Save stores user's reading progress.
func (h *ProgressHandler) Save(c *gin.Context) {
	var req progressSaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uint(userIDVal.(float64))
	if err := h.svc.Save(userID, req.NovelID, req.ChapterID, req.PositionOffset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Get returns user's progress for a novel.
func (h *ProgressHandler) Get(c *gin.Context) {
	novelIDParam := c.Param("novelId")
	novelID64, err := strconv.ParseUint(novelIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid novel id"})
		return
	}
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uint(userIDVal.(float64))
	rp, err := h.svc.Get(userID, uint(novelID64))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"novelId":        rp.NovelID,
		"chapterId":      rp.ChapterID,
		"positionOffset": rp.Position,
	})
}
