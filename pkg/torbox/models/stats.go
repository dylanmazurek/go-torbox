package models

// Stats represents TorBox usage statistics
type Stats struct {
	TotalDownloaded int64 `json:"total_downloaded"`
	TotalUploaded   int64 `json:"total_uploaded"`
	ActiveTorrents  int   `json:"active_torrents"`
	QueuedTorrents  int   `json:"queued_torrents"`
	TotalTorrents   int   `json:"total_torrents"`
}

// StatsResponse represents the response from the stats endpoint
type StatsResponse struct {
	BaseResponse
	Data *Stats `json:"data"`
}
