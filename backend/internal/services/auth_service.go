package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/truyentan/backend/internal/models"
)

// AuthService handles user authentication and JWT generation.
type AuthService struct {
	db         *gorm.DB
	jwtSecret  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// NewAuthService creates a new AuthService.
func NewAuthService(db *gorm.DB, secret string, accessTTL, refreshTTL time.Duration) *AuthService {
	return &AuthService{db: db, jwtSecret: secret, accessTTL: accessTTL, refreshTTL: refreshTTL}
}

// TokenPair represents generated access and refresh tokens.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Signup creates a new user and returns JWT tokens.
func (s *AuthService) Signup(name, email, password string) (*TokenPair, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{Name: name, Email: email, PasswordHash: string(hash)}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return s.generateTokens(user.ID)
}

// Signin authenticates a user and returns JWT tokens.
func (s *AuthService) Signin(email, password string) (*TokenPair, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}
	return s.generateTokens(user.ID)
}

// Refresh generates new tokens using a refresh token.
func (s *AuthService) Refresh(refreshToken string) (*TokenPair, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	sub, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid token subject")
	}
	return s.generateTokens(uint(sub))
}

func (s *AuthService) generateTokens(userID uint) (*TokenPair, error) {
	access, err := s.generateToken(userID, s.accessTTL)
	if err != nil {
		return nil, err
	}
	refresh, err := s.generateToken(userID, s.refreshTTL)
	if err != nil {
		return nil, err
	}
	return &TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *AuthService) generateToken(userID uint, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
