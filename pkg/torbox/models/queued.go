package models

type QueuedDownload struct {
	ID          int64   `json:"id"`
	CreatedAt   string  `json:"created_at"`
	Magnet      *string `json:"magnet"`
	TorrentFile *string `json:"torrent_file"`
	Hash        string  `json:"hash"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
}
