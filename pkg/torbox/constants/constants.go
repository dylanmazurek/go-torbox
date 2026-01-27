package constants

const (
	API_GENERAL_BASE_URL = "https://api.torbox.app/v1"
	API_SEARCH_BASE_URL  = "https://search-api.torbox.app"
)

// Torrent API Paths
const (
	PATH_TORRENTS_GET_ACTIVE       = "api/torrents/mylist"
	PATH_TORRENTS_CREATE           = "api/torrents/createtorrent"
	PATH_TORRENTS_CONTROL_ACTIVE   = "api/torrents/controltorrent"
	PATH_TORRENTS_INFO             = "api/torrents/torrentinfo"
	PATH_TORRENTS_GET_DOWNLOAD_URL = "api/torrents/requestdl"
	PATH_TORRENTS_CHECK_CACHED     = "api/torrents/checkcached"
	PATH_TORRENTS_EXPORT_DATA      = "api/torrents/exportdata"
	PATH_TORRENTS_MAGNET_TO_FILE   = "api/torrents/magnettofile"
)

// Queued Torrent API Paths
const (
	PATH_TORRENTS_GET_QUEUED   = "api/queued/getqueued"
	PATH_TORRENTS_CONTROL_QUEUED = "api/queued/controlqueued"
)

// Usenet API Paths
const (
	PATH_USENET_CREATE       = "api/usenet/createusenetdownload"
	PATH_USENET_GET_LIST     = "api/usenet/mylist"
	PATH_USENET_CONTROL      = "api/usenet/controlusenetdownload"
	PATH_USENET_GET_DOWNLOAD = "api/usenet/requestdl"
	PATH_USENET_CHECK_CACHED = "api/usenet/checkcached"
)

// Web Download API Paths
const (
	PATH_WEBDL_CREATE  = "api/webdl/createwebdownload"
	PATH_WEBDL_GET_LIST = "api/webdl/mylist"
	PATH_WEBDL_CONTROL = "api/webdl/controlwebdownload"
	PATH_WEBDL_GET_DOWNLOAD = "api/webdl/requestdl"
	PATH_WEBDL_CHECK_CACHED = "api/webdl/checkcached"
)

// User API Paths
const (
	PATH_USER_ME            = "api/user/me"
	PATH_USER_REFRESH_TOKEN = "api/user/refreshtoken"
	PATH_USER_ADD_REFERRAL  = "api/user/addreferral"
)

// Stats API Paths
const (
	PATH_STATS = "api/stats"
)

// Notifications API Paths
const (
	PATH_NOTIFICATIONS_LIST  = "api/notifications/mynotifications"
	PATH_NOTIFICATIONS_CLEAR = "api/notifications/clear"
	PATH_NOTIFICATIONS_RSS   = "api/notifications/rss"
)

// RSS API Paths
const (
	PATH_RSS_ADD_FEED    = "api/rss/addrss"
	PATH_RSS_MODIFY_FEED = "api/rss/modifyrss"
	PATH_RSS_CONTROL     = "api/rss/controlrss"
	PATH_RSS_GET_FEEDS   = "api/rss/getfeeds"
	PATH_RSS_GET_ITEMS   = "api/rss/getfeeditems"
)

// Search API Paths
const (
	PATH_SEARCH_TORRENTS = "torrents"
	PATH_SEARCH_META     = "meta"
	PATH_TORRENTS_SEARCH = "api/torrents/search"
	PATH_TORRENTS_STORE_SEARCH = "api/torrents/storesearch"
)
