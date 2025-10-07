package constants

const (
	API_GENERAL_BASE_URL = "https://api.torbox.app/v1"
	API_SEARCH_BASE_URL  = "https://search-api.torbox.app"
)

// GENERAL API
const (
	PATH_TORRENTS_GET_ACTIVE       = "api/torrents/mylist"
	PATH_TORRENTS_GET_DOWNLOAD_URL = "api/torrents/requestdl"
	PATH_TORRENTS_CREATE           = "api/torrents/createtorrent"
	PATH_TORRENTS_CONTROL_ACTIVE   = "api/torrents/controltorrent"
	PATH_TORRENTS_CHECK_CACHED     = "api/torrents/checkcached"
	PATH_TORRENTS_INFO             = "api/torrents/torrentinfo"
	PATH_TORRENTS_EXPORT_DATA      = "api/torrents/exportdata"
	PATH_TORRENTS_SEARCH           = "api/torrents/search"
	PATH_TORRENTS_STORE_SEARCH     = "api/torrents/storesearch"

	PATH_TORRENTS_GET_QUEUED     = "api/queued/getqueued"
	PATH_TORRENTS_CONTROL_QUEUED = "api/queued/controlqueued"

	// Usenet API
	PATH_USENET_CREATE         = "api/usenet/createusenetdownload"
	PATH_USENET_CONTROL        = "api/usenet/controlusenetdownload"
	PATH_USENET_GET_DOWNLOAD   = "api/usenet/requestdl"
	PATH_USENET_GET_LIST       = "api/usenet/mylist"
	PATH_USENET_CHECK_CACHED   = "api/usenet/checkcached"

	// Web Downloads API
	PATH_WEBDL_CREATE  = "api/webdl/createwebdownload"
	PATH_WEBDL_CONTROL = "api/webdl/controlwebdownload"

	// User API
	PATH_USER_ME            = "api/user/me"
	PATH_USER_REFRESH_TOKEN = "api/user/refreshtoken"
	PATH_USER_ADD_REFERRAL  = "api/user/addreferral"

	// Notifications API
	PATH_NOTIFICATIONS_RSS   = "api/notifications/rss"
	PATH_NOTIFICATIONS_LIST  = "api/notifications/mynotifications"
	PATH_NOTIFICATIONS_CLEAR = "api/notifications/clear"

	// RSS API
	PATH_RSS_ADD     = "api/rss/addrss"
	PATH_RSS_CONTROL = "api/rss/controlrss"
	PATH_RSS_MODIFY  = "api/rss/modifyrss"

	// Integration API
	PATH_INTEGRATION_GOOGLEDRIVE = "api/integration/googledrive"
	PATH_INTEGRATION_DROPBOX     = "api/integration/dropbox"
	PATH_INTEGRATION_ONEDRIVE    = "api/integration/onedrive"
	PATH_INTEGRATION_GOFILE      = "api/integration/gofile"
	PATH_INTEGRATION_1FICHIER    = "api/integration/1fichier"
	PATH_INTEGRATION_JOBS        = "api/integration/jobs"

	// Stats API
	PATH_STATS = "api/stats"
)

// SEARCH API
const (
	PATH_SEARCH_TORRENTS = "torrents"
	PATH_SEARCH_META     = "meta"
)

type SeedSetting int

const (
	Auto SeedSetting = iota + 1
	Seed
	NoSeed
)
