package notification

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

func (h *Handler) GetNotifications(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	notifs, err := h.service.GetNotifications(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "FETCH_FAILED"})
		return
	}

	c.JSON(http.StatusOK, notifs)
}

func (h *Handler) MarkRead(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "code": "UNAUTHORIZED"})
		return
	}

	notifID := c.Param("id")
	err := h.service.MarkRead(c.Request.Context(), userID.(string), notifID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": "UPDATE_FAILED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification marked as read"})
}
