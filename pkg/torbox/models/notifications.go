package models

// Notification represents a user notification
type Notification struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"created_at"`
}

// NotificationsResponse represents the response from the notifications list endpoint
type NotificationsResponse struct {
	BaseResponse
	Data []Notification `json:"data"`
}

// ClearNotificationRequest represents a request to clear a notification
type ClearNotificationRequest struct {
	ID *int64 `json:"id,omitempty"` // If nil, clear all notifications
}
