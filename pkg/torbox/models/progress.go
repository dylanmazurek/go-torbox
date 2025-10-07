package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

type ProgressDetails struct {
	DownloadPresent  bool                   `json:"download_present"`
	DownloadPath     string                 `json:"download_path"`
	DownloadFinished bool                   `json:"download_finished"`
	TotalUploaded    int64                  `json:"total_uploaded"`
	TotalDownloaded  int64                  `json:"total_downloaded"`
	DownloadState    constants.TorrentState `json:"download_state"`
	Progress         float64                `json:"progress"`
	DownloadSpeed    int64                  `json:"download_speed"`
	UploadSpeed      int64                  `json:"upload_speed"`
}
