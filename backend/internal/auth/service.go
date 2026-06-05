package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidToken       = errors.New("invalid token")
)

type Service interface {
	Register(ctx context.Context, name, email, password, userType string, income float64) (*User, error)
	Login(ctx context.Context, email, password string) (*User, string, string, error)
	Refresh(ctx context.Context, refreshTokenStr string) (string, string, error)
	Logout(ctx context.Context, token string, refreshTokenStr string) error
	GetUserProfile(ctx context.Context, userID string) (*User, error)
	UpdateUserProfile(ctx context.Context, userID string, name string, income float64, userType string) (*User, error)
}

type service struct {
	repo             Repository
	jwtSecret        []byte
	jwtRefreshSecret []byte
}

func NewService(repo Repository, jwtSecret, jwtRefreshSecret string) Service {
	return &service{
		repo:             repo,
		jwtSecret:        []byte(jwtSecret),
		jwtRefreshSecret: []byte(jwtRefreshSecret),
	}
}

func (s *service) Register(ctx context.Context, name, email, password, userType string, income float64) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		UserType: userType,
		Income:   income,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Login(ctx context.Context, email, password string) (*User, string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, "", "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", ErrInvalidCredentials
	}

	accessToken, err := s.generateAccessToken(user.ID)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := s.createRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, refreshTokenStr string) (string, string, error) {
	// Replay detection, rotation, etc. will go here.
	return "", "", nil
}

func (s *service) Logout(ctx context.Context, token string, refreshTokenStr string) error {
	_ = s.repo.DeleteSession(ctx, token)
	_ = s.repo.DeleteRefreshToken(ctx, refreshTokenStr)
	return nil
}

func (s *service) GetUserProfile(ctx context.Context, userID string) (*User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *service) UpdateUserProfile(ctx context.Context, userID string, name string, income float64, userType string) (*User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if income > 0 {
		user.Income = income
	}
	if userType != "" {
		user.UserType = userType
	}

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) generateAccessToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})
	return token.SignedString(s.jwtSecret)
}

func (s *service) createRefreshToken(ctx context.Context, userID string) (string, error) {
	// Simple mock token generation
	tokenStr := "opaque_refresh_" + userID + "_" + time.Now().Format("20060102150405")
	token := &RefreshToken{
		Token:     tokenStr,
		UserID:    userID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}
	if err := s.repo.CreateRefreshToken(ctx, token); err != nil {
		return "", err
	}
	return tokenStr, nil
}
