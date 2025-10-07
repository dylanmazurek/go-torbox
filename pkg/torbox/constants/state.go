package constants

type TorrentState string

const (
	// ---- processing states
	// checkingResumeData	The torrent is checking resumable data.
	TorrentStateCheckingResumeData TorrentState = "checkingResumeData"
	// checking
	TorrentStateChecking = "checking"
	// metaDL			The torrent is downloading metadata from the hoard.
	TorrentStateMetaDL = "metaDL"
	// paused			The torrent is paused.
	TorrentStatePaused = "paused"

	// ---- download states
	// downloading		The torrent is currently downloading.
	TorrentStateDownloading = "downloading"
	// stalled (no seeds)	The torrent is trying to download, but there are no seeds connected to download from.
	TorrentStateStalledNoSeeds = "stalled (no seeds)"
	// stalled
	TorrentStateStalledDL = "stalledDL"

	// ---- upload states
	// uploading		The torrent is currently seeding.
	TorrentStateUploading = "uploading"
	// uploading (no peers) The torrent is currently seeding, but there are no peers connected to upload to.
	TorrentStateUploadingNoPeers = "uploading (no peers)"

	// ---- completion states
	// completed		The torrent is completely downloaded. Do not use this for download completion status.
	TorrentStateCompleted = "completed"
	// cached			The torrent is cached from the server.
	TorrentStateCached = "cached"

	// ---- error states
	// unknown			The torrent state is unknown.
	TorrentStateUnknown = "unknown"
)

func (t TorrentState) IsComplete() bool {
	switch t {
	case TorrentStateCompleted,
		TorrentStateCached,
		TorrentStateUploading,
		TorrentStateUploadingNoPeers:

		return true
	default:
		return false
	}
}
