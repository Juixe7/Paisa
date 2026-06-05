package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type manualTxRequest struct {
	Amount       float64 `json:"amount" binding:"required,gt=0"`
	MerchantName string  `json:"merchant_name" binding:"required"`
	CategoryID   string  `json:"category_id" binding:"required"`
}

func (h *Handler) CreateManual(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	var req manualTxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	tx, err := h.service.CreateManualTransaction(c.Request.Context(), userID.(string), req.Amount, req.MerchantName, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "TRANSACTION_FAILED"})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func (h *Handler) List(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// Build basic filters map
	filters := make(map[string]interface{})
	filters["user_id"] = userID.(string)
	if catID := c.Query("category_id"); catID != "" {
		filters["category_id"] = catID
	}

	// Scaffold lists (empty mock list)
	c.JSON(http.StatusOK, gin.H{
		"transactions": []interface{}{},
		"limit":        limit,
		"offset":       offset,
	})
}

func (h *Handler) Detail(c *gin.Context) {
	txID := c.Param("id")
	// Scaffold detail (empty mock)
	c.JSON(http.StatusOK, gin.H{
		"id":            txID,
		"merchant_name": "Mock Merchant",
		"amount":        100.0,
		"direction":     "debit",
	})
}

type smsRequest struct {
	RawSMS string `json:"raw_sms" binding:"required"`
}

func (h *Handler) IngestSMS(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	var req smsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	tx, err := h.service.IngestSMS(c.Request.Context(), userID.(string), req.RawSMS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "INGEST_FAILED"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

type intentRequest struct {
	Amount       float64 `json:"amount" binding:"required,gt=0"`
	CategoryID   string  `json:"category_id" binding:"required"`
	MerchantName string  `json:"merchant_name" binding:"required"`
}

func (h *Handler) CreateIntent(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	var req intentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "INVALID_INPUT"})
		return
	}

	intent, err := h.service.CreatePrePaymentIntent(c.Request.Context(), userID.(string), req.Amount, req.CategoryID, req.MerchantName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "INTENT_FAILED"})
		return
	}

	c.JSON(http.StatusCreated, intent)
}

func (h *Handler) ConfirmIntent(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	intentID := c.Param("id")
	tx, err := h.service.ConfirmPrePaymentIntent(c.Request.Context(), userID.(string), intentID)
	if err != nil {
		if err == ErrIntentExpired {
			c.JSON(http.StatusGone, gin.H{"error": err.Error(), "code": "INTENT_EXPIRED"})
			return
		}
		if err == ErrIntentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "code": "INTENT_NOT_FOUND"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "CONFIRM_FAILED"})
		return
	}

	c.JSON(http.StatusOK, tx)
}
