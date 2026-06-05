package analytics

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

func (h *Handler) GetSummary(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	summary, err := h.service.GetSummary(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "FETCH_FAILED"})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (h *Handler) GetBudgets(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	budgets, err := h.service.GetBudgets(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "FETCH_FAILED"})
		return
	}

	c.JSON(http.StatusOK, budgets)
}

type updateBudgetRequest struct {
	Limit float64 `json:"limit" binding:"required,gt=0"`
}

func (h *Handler) UpdateBudget(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	categoryID := c.Param("category_id")

	var req updateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	err := h.service.UpdateBudget(c.Request.Context(), userID.(string), categoryID, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "UPDATE_FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "budget limit updated successfully"})
}
