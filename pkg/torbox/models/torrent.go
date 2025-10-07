package models

import (
	"encoding/json"
	"time"
)

type Torrent struct {
	TrackerDetails
	ProgressDetails

	ID     int64  `json:"id"`
	Hash   string `json:"hash"`
	Server int64  `json:"server"`

	AuthID      string `json:"auth_id"`
	Name        string `json:"name"`
	Magnet      string `json:"magnet"`
	Size        int64  `json:"size"`
	ETA         int64  `json:"eta"`
	TorrentFile bool   `json:"torrent_file"`

	Cached      bool   `json:"cached"`
	Owner       string `json:"owner"`
	AllowZipped bool   `json:"allow_zipped"`
	ShortName   string `json:"short_name"`

	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	ExpiresAt *time.Time `json:"-"`

	Files []File `json:"-"`
}

func (t *Torrent) IsDownloaded() bool {
	return t.DownloadFinished
}

func (t *Torrent) UnmarshalJSON(d []byte) error {
	type Alias Torrent
	type Aux struct {
		*Alias

		CreatedAtStr string `json:"created_at"`
		UpdatedAtStr string `json:"updated_at"`
		ExpiresAtStr string `json:"expires_at"`

		TorrentID *int64 `json:"torrent_id"`
		QueuedID  *int64 `json:"queued_id"`
		Files     any    `json:"files"`
	}

	aux := &Aux{
		Alias: (*Alias)(t),
	}

	err := json.Unmarshal(d, &aux)
	if err != nil {
		return err
	}

	if aux.CreatedAtStr != "" {
		createdAt, err := time.Parse("2006-01-02T15:04:05Z", aux.CreatedAtStr)
		if err != nil {
			return err
		}

		t.CreatedAt = &createdAt
	}

	if aux.UpdatedAtStr != "" {
		updatedAt, err := time.Parse("2006-01-02T15:04:05Z", aux.UpdatedAtStr)
		if err != nil {
			return err
		}

		t.UpdatedAt = &updatedAt
	}

	if aux.ExpiresAtStr != "" {
		expiresAt, err := time.Parse("2006-01-02T15:04:05Z", aux.ExpiresAtStr)
		if err != nil {
			return err
		}

		t.ExpiresAt = &expiresAt
	}

	if aux.TorrentID != nil {
		t.ID = *aux.TorrentID
	}

	if aux.QueuedID != nil {
		t.ID = *aux.QueuedID
	}

	if aux.Files == nil {
		return nil
	}

	return err
}
