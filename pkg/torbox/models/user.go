package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

// User models
type User struct {
	ID            int64              `json:"id"`
	Email         string             `json:"email"`
	Plan          constants.PlanType `json:"plan"`
	PremiumExpiry string             `json:"premium_expiry"`
	CooldownUntil string             `json:"cooldown_until"`
	Customer      string             `json:"customer"`
	Server        int                `json:"server"`
	IsSubscribed  bool               `json:"is_subscribed"`
	UserReferral  string             `json:"user_referral"`

	Auth0ID   *string `json:"auth0_id,omitempty"`
	BaseEmail *string `json:"base_email,omitempty"`

	TotalDownloaded           int64  `json:"total_downloaded"`
	TotalUploaded             int64  `json:"total_uploaded"`
	TotalBytesDownloaded      int64  `json:"total_bytes_downloaded"`
	TotalBytesUploaded        int64  `json:"total_bytes_uploaded"`
	TorrentsDownloaded        int64  `json:"torrents_downloaded"`
	WebDownloadsDownloaded    int64  `json:"web_downloads_downloaded"`
	UsenetDownloadsDownloaded int64  `json:"usenet_downloads_downloaded"`
	AdditionalConcurrentSlots int64  `json:"additional_concurrent_slots"`
	LongTermSeeding           bool   `json:"long_term_seeding"`
	LongTermStorage           bool   `json:"long_term_storage"`
	IsVendor                  bool   `json:"is_vendor"`
	VendorID                  *int64 `json:"vendor_id,omitempty"`
	PurchasesReferred         int64  `json:"purchases_referred"`
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
