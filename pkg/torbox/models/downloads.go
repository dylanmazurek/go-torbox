package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

// Usenet models
type UsenetDownload struct {
	ID             int64   `json:"id"`
	Hash           string  `json:"hash"`
	Name           string  `json:"name"`
	Size           int64   `json:"size"`
	DownloadState  string  `json:"download_state"`
	DownloadSpeed  float64 `json:"download_speed"`
	UploadSpeed    float64 `json:"upload_speed"`
	DownloadedSize int64   `json:"downloaded"`
	Progress       float64 `json:"progress"`
	Ratio          float64 `json:"ratio"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Files          []File  `json:"files"`
}

type CreateUsenetRequest struct {
	Link     string  `json:"link"`
	Name     *string `json:"name,omitempty"`
	AsQueued *bool   `json:"as_queued,omitempty"`
}

type CreateUsenetResponse struct {
	BaseResponse
	Data *UsenetDownload `json:"data"`
}

type GetUsenetListResponse struct {
	BaseResponse
	Data []UsenetDownload `json:"data"`
}

type ControlUsenetRequest struct {
	UsenetID  int64                          `json:"usenet_id"`
	Operation constants.ControlUsenetOperation `json:"operation"`
}

// Web Download models
type WebDownload struct {
	ID             int64   `json:"id"`
	Hash           string  `json:"hash"`
	Name           string  `json:"name"`
	Size           int64   `json:"size"`
	DownloadState  string  `json:"download_state"`
	DownloadSpeed  float64 `json:"download_speed"`
	DownloadedSize int64   `json:"downloaded"`
	Progress       float64 `json:"progress"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Files          []File  `json:"files"`
}

type CreateWebDownloadRequest struct {
	Link     string  `json:"link"`
	Name     *string `json:"name,omitempty"`
	AsQueued *bool   `json:"as_queued,omitempty"`
}

type CreateWebDownloadResponse struct {
	BaseResponse
	Data *WebDownload `json:"data"`
}

type ControlWebDownloadRequest struct {
	WebID     int64                                 `json:"web_id"`
	Operation constants.ControlWebDownloadOperation `json:"operation"`
}
