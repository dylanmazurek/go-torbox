package models

type GetMetaResponse struct {
	BaseResponse
	Data Torrent `json:"data"`
}

type GetTorrentResponse struct {
	BaseResponse
	Data struct {
		Metadata Metadata  `json:"metadata"`
		Torrents []Torrent `json:"torrents"`
	} `json:"data"`
}

type CheckCachedResponse struct {
	BaseResponse
	Data *CacheCheckResponse `json:"data"`
}

type CacheCheckResponse struct {
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Hash   string `json:"hash"`
	Cached bool   `json:"cached"`
}
