package goals

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type createGoalRequest struct {
	Title        string    `json:"title" binding:"required"`
	TargetAmount float64   `json:"target_amount" binding:"required,gt=0"`
	TargetDate   time.Time `json:"target_date" binding:"required"`
}

func (h *Handler) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	var req createGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	goal, err := h.service.CreateGoal(c.Request.Context(), userID.(string), req.Title, req.TargetAmount, req.TargetDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "GOAL_FAILED"})
		return
	}

	c.JSON(http.StatusCreated, goal)
}

func (h *Handler) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	goals, err := h.service.GetGoals(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "FETCH_FAILED"})
		return
	}

	c.JSON(http.StatusOK, goals)
}

type contributeRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

func (h *Handler) Contribute(c *gin.Context) {
	goalID := c.Param("id")

	var req contributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	err := h.service.Contribute(c.Request.Context(), goalID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "CONTRIBUTION_FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contribution successful"})
}
