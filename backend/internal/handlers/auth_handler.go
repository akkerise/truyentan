package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/truyentan/backend/internal/services"
)

// AuthHandler provides HTTP handlers for authentication endpoints.
type AuthHandler struct {
	svc *services.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(svc *services.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// signupRequest represents signup payload.
type signupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// signinRequest represents signin payload.
type signinRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// refreshRequest represents token refresh payload.
type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Signup godoc
// @Summary User signup
// @Description Create new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body signupRequest true "signup request"
// @Success 201 {object} services.TokenPair
// @Failure 400 {object} gin.H
// @Router /api/v1/auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.svc.Signup(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tokens)
}

// Signin godoc
// @Summary User signin
// @Tags auth
// @Accept json
// @Produce json
// @Param request body signinRequest true "signin request"
// @Success 200 {object} services.TokenPair
// @Failure 401 {object} gin.H
// @Router /api/v1/auth/signin [post]
func (h *AuthHandler) Signin(c *gin.Context) {
	var req signinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.svc.Signin(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}

// Refresh godoc
// @Summary Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body refreshRequest true "refresh request"
// @Success 200 {object} services.TokenPair
// @Failure 401 {object} gin.H
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.svc.Refresh(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}
