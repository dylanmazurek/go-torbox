package constants

const (
	// Torrents API
	PATH_TORRENTS_GET_ACTIVE       = "api/torrents/mylist"
	PATH_TORRENTS_GET_DOWNLOAD_URL = "api/torrents/requestdl"
	PATH_TORRENTS_CREATE           = "api/torrents/createtorrent"
	PATH_TORRENTS_CONTROL_ACTIVE   = "api/torrents/controltorrent"
	PATH_TORRENTS_CHECK_CACHED     = "api/torrents/checkcached"
	PATH_TORRENTS_INFO             = "api/torrents/torrentinfo"
	PATH_TORRENTS_EXPORT_DATA      = "api/torrents/exportdata"
	PATH_TORRENTS_SEARCH           = "api/torrents/search"
	PATH_TORRENTS_STORE_SEARCH     = "api/torrents/storesearch"

	// Queued Torrents API
	PATH_TORRENTS_GET_QUEUED     = "api/queued/getqueued"
	PATH_TORRENTS_CONTROL_QUEUED = "api/queued/controlqueued"

	// Usenet API
	PATH_USENET_CREATE       = "api/usenet/createusenetdownload"
	PATH_USENET_CONTROL      = "api/usenet/controlusenetdownload"
	PATH_USENET_GET_DOWNLOAD = "api/usenet/requestdl"
	PATH_USENET_GET_LIST     = "api/usenet/mylist"
	PATH_USENET_CHECK_CACHED = "api/usenet/checkcached"

	// Web Downloads API
	PATH_WEBDL_CREATE  = "api/webdl/createwebdownload"
	PATH_WEBDL_CONTROL = "api/webdl/controlwebdownload"

	// User API
	PATH_USER_ME            = "api/user/me"
	PATH_USER_REFRESH_TOKEN = "api/user/refreshtoken"
	PATH_USER_ADD_REFERRAL  = "api/user/addreferral"

	// Stats API
	PATH_STATS = "api/stats"
)
