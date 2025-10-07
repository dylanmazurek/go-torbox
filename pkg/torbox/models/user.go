package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

// User models
type User struct {
	ID              int64   `json:"id"`
	Email           string  `json:"email"`
	Plan            string  `json:"plan"`
	PremiumExpiry   string  `json:"premium_expiry"`
	CooldownUntil   string  `json:"cooldown_until"`
	Auth0ID         string  `json:"auth0_id"`
	TotalDownloaded int64   `json:"total_downloaded"`
	TotalUploaded   int64   `json:"total_uploaded"`
	Customer        string  `json:"customer"`
	Server          int     `json:"server"`
	IsSubscribed    bool    `json:"is_subscribed"`
	UserReferral    string  `json:"user_referral"`
	BaseEmail       *string `json:"base_email,omitempty"`
}

type GetUserResponse struct {
	BaseResponse
	Data *User `json:"data"`
}

type RefreshTokenResponse struct {
	BaseResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type AddReferralRequest struct {
	ReferralCode string `json:"referral_code"`
}

// Notification models
type Notification struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"created_at"`
}

type GetNotificationsResponse struct {
	BaseResponse
	Data []Notification `json:"data"`
}

// RSS models
type RSSFeed struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddRSSRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type AddRSSResponse struct {
	BaseResponse
	Data *RSSFeed `json:"data"`
}

type ControlRSSRequest struct {
	RSSID     int64                       `json:"rss_id"`
	Operation constants.ControlRSSOperation `json:"operation"`
}

type ModifyRSSRequest struct {
	RSSID   int64   `json:"rss_id"`
	URL     *string `json:"url,omitempty"`
	Name    *string `json:"name,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

type ModifyRSSResponse struct {
	BaseResponse
	Data *RSSFeed `json:"data"`
}

// Integration models
type IntegrationAuthRequest struct {
	Code   string `json:"code,omitempty"`
	APIKey string `json:"api_key,omitempty"`
}

type IntegrationJob struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size"`
	Progress    float64 `json:"progress"`
	Destination string `json:"destination"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetIntegrationJobsResponse struct {
	BaseResponse
	Data []IntegrationJob `json:"data"`
}

// Stats models
type Stats struct {
	TotalDownloaded int64   `json:"total_downloaded"`
	TotalUploaded   int64   `json:"total_uploaded"`
	TotalTorrents   int     `json:"total_torrents"`
	ActiveTorrents  int     `json:"active_torrents"`
	QueuedTorrents  int     `json:"queued_torrents"`
	TotalUsenet     int     `json:"total_usenet"`
	TotalWebDL      int     `json:"total_webdl"`
	AvailableSpace  int64   `json:"available_space"`
	UsedSpace       int64   `json:"used_space"`
	Plan            string  `json:"plan"`
	PremiumExpiry   *string `json:"premium_expiry,omitempty"`
}

type GetStatsResponse struct {
	BaseResponse
	Data *Stats `json:"data"`
}
