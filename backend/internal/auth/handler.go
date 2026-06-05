package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type registerRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	UserType string  `json:"user_type" binding:"required,oneof=student professional"`
	Income   float64 `json:"income" binding:"required,gte=0"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	user, err := h.service.Register(c.Request.Context(), req.Name, req.Email, req.Password, req.UserType, req.Income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "REGISTRATION_FAILED"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	user, access, refresh, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "code": "INVALID_CREDENTIALS"})
		return
	}

	// Set refresh token in HttpOnly cookie
	c.SetCookie("refresh_token", refresh, 30*24*3600, "/auth/refresh", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user":          user,
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	cookieToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token", "code": "MISSING_TOKEN"})
		return
	}

	access, refresh, err := h.service.Refresh(c.Request.Context(), cookieToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "code": "REFRESH_FAILED"})
		return
	}

	c.SetCookie("refresh_token", refresh, 30*24*3600, "/auth/refresh", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	cookieToken, _ := c.Cookie("refresh_token")

	// Call logout service
	_ = h.service.Logout(c.Request.Context(), authHeader, cookieToken)

	c.SetCookie("refresh_token", "", -1, "/auth/refresh", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}

func (h *Handler) GetMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	user, err := h.service.GetUserProfile(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "code": "USER_NOT_FOUND"})
		return
	}

	c.JSON(http.StatusOK, user)
}

type updateProfileRequest struct {
	Name     string  `json:"name"`
	Income   float64 `json:"income" binding:"omitempty,gte=0"`
	UserType string  `json:"user_type" binding:"omitempty,oneof=student professional"`
}

func (h *Handler) UpdateMe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	var req updateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	user, err := h.service.UpdateUserProfile(c.Request.Context(), userID.(string), req.Name, req.Income, req.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "UPDATE_FAILED"})
		return
	}

	c.JSON(http.StatusOK, user)
}
