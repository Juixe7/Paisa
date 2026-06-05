package categorisation

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

type correctionRequest struct {
	OldCategoryID string `json:"old_category_id" binding:"required"`
	NewCategoryID string `json:"new_category_id" binding:"required"`
	SubcategoryID string `json:"subcategory_id"`
	MerchantName  string `json:"merchant_name" binding:"required"`
}

func (h *Handler) CorrectCategory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	txID := c.Param("id")

	var req correctionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	err := h.service.CorrectCategory(
		c.Request.Context(),
		userID.(string),
		txID,
		req.OldCategoryID,
		req.NewCategoryID,
		req.SubcategoryID,
		req.MerchantName,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "CORRECTION_FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category corrected successfully"})
}

func (h *Handler) GetCategories(c *gin.Context) {
	cats, subcats, err := h.service.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "FETCH_FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories":    cats,
		"subcategories": subcats,
	})
}
