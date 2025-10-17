package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotificationHandler handles notification-related HTTP requests
type NotificationHandler struct {
	// TODO: Add use case dependencies
}

// NewNotificationHandler creates a new NotificationHandler
func NewNotificationHandler(useCase interface{}) *NotificationHandler {
	return &NotificationHandler{}
}

// List lists all notifications for the authenticated user
// @Summary List notifications
// @Description Get all notifications for the current user
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /notifications [get]
func (h *NotificationHandler) List(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "List notifications - TODO"})
}

// ListUnread lists unread notifications
// @Summary List unread notifications
// @Description Get all unread notifications for the current user
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /notifications/unread [get]
func (h *NotificationHandler) ListUnread(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "List unread notifications - TODO"})
}

// MarkAsRead marks a notification as read
// @Summary Mark notification as read
// @Description Mark a specific notification as read
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Notification ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /notifications/{id}/read [put]
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Mark as read - TODO"})
}

// MarkAllAsRead marks all notifications as read
// @Summary Mark all notifications as read
// @Description Mark all notifications as read for the current user
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /notifications/read-all [put]
func (h *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Mark all as read - TODO"})
}
