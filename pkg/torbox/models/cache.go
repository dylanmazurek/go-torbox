package models

// CheckCachedTorrentsRequest represents a request to check if torrents are cached
type CheckCachedTorrentsRequest struct {
	Hashes []string `json:"hashes"`
}

// MagnetToFileRequest represents a request to convert a magnet link to a torrent file
type MagnetToFileRequest struct {
	Magnet string `json:"magnet"`
}

// MagnetToFileResponse represents the response from converting a magnet to file
type MagnetToFileResponse struct {
	BaseResponse
	Data string `json:"data"` // Base64 encoded torrent file
}

// ExportDataRequest represents a request to export torrent data
type ExportDataRequest struct {
	TorrentID int64  `json:"torrent_id"`
	Type      string `json:"type"` // "json" or "csv"
}
