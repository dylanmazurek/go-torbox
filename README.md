# go-torbox

A comprehensive Go SDK for the [TorBox API](https://torbox.app), providing easy access to torrent, usenet, web downloads, and various cloud integrations.

## Installation

```bash
go get github.com/dylanmazurek/go-torbox
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/dylanmazurek/go-torbox/pkg/torbox"
)

func main() {
    ctx := context.Background()
    
    client, err := torbox.New(ctx, torbox.WithAPIKey("your-api-key"))
    if err != nil {
        log.Fatal(err)
    }

    // Get active torrents
    torrents, err := client.General.GetActiveTorrents()
    if err != nil {
        log.Fatal(err)
    }

    for _, t := range torrents {
        fmt.Printf("Torrent: %s (%.2f%%)\n", t.Name, t.Progress*100)
    }
}
```

## Features

### Torrent Management
- ✅ Create torrents from magnet links or .torrent files
- ✅ List active and queued torrents
- ✅ Control torrents (pause, resume, delete, reannounce)
- ✅ Get download URLs for torrent files
- ✅ Check if torrents are cached
- ✅ Get torrent info by hash
- ✅ Search torrents
- ✅ Export torrent data

### Usenet Downloads
- ✅ Create usenet downloads
- ✅ List usenet downloads
- ✅ Control usenet downloads (pause, resume, delete)
- ✅ Get download URLs
- ✅ Check usenet cache

### Web Downloads
- ✅ Create web downloads from HTTP/HTTPS URLs
- ✅ Control web downloads (pause, resume, delete)

### User Management
- ✅ Get user information
- ✅ Refresh authentication token
- ✅ Add referral codes

### Notifications
- ✅ Get RSS notifications
- ✅ Get all notifications
- ✅ Clear notifications

### RSS Feeds
- ✅ Add RSS feeds
- ✅ Control RSS feeds (pause, resume, delete)
- ✅ Modify RSS feed settings

### Cloud Integrations
- ✅ Google Drive authorization
- ✅ Dropbox authorization
- ✅ OneDrive authorization
- ✅ Gofile integration
- ✅ 1Fichier integration
- ✅ Get integration job status

### Statistics
- ✅ Get account statistics and usage

### Search API
- ✅ Search torrents by metadata
- ✅ Get torrent metadata

## Usage Examples

### Creating a Torrent from Magnet Link

```go
import (
    "github.com/dylanmazurek/go-torbox/pkg/magnet"
    "github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

magnetLink := "magnet:?xt=urn:btih:..."
mag, err := magnet.NewMagnet(magnetLink)
if err != nil {
    log.Fatal(err)
}

name := "My Download"
req := models.CreateTorrentRequest{
    Magnet: mag,
    Name:   &name,
}

torrent, err := client.General.CreateTorrent(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created torrent: %s (ID: %d)\n", torrent.Name, torrent.ID)
```

### Checking if a Torrent is Cached

```go
hash := "abc123def456..."
cacheInfo, err := client.General.CheckCached(hash)
if err != nil {
    log.Fatal(err)
}

if cacheInfo.Cached {
    fmt.Printf("Torrent is cached! Name: %s, Size: %d\n", cacheInfo.Name, cacheInfo.Size)
} else {
    fmt.Println("Torrent is not cached")
}
```

### Creating a Usenet Download

```go
link := "nzb-link-or-url"
name := "My Usenet Download"
req := models.CreateUsenetRequest{
    Link: link,
    Name: &name,
}

usenet, err := client.General.CreateUsenetDownload(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created usenet download: %s\n", usenet.Name)
```

### Creating a Web Download

```go
link := "https://example.com/file.zip"
name := "My Web Download"
req := models.CreateWebDownloadRequest{
    Link: link,
    Name: &name,
}

webdl, err := client.General.CreateWebDownload(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created web download: %s\n", webdl.Name)
```

### Getting User Information

```go
user, err := client.General.GetUser()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("User: %s\n", user.Email)
fmt.Printf("Plan: %s\n", user.Plan)
fmt.Printf("Downloaded: %d bytes\n", user.TotalDownloaded)
```

### Getting Account Statistics

```go
stats, err := client.General.GetStats()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total Torrents: %d\n", stats.TotalTorrents)
fmt.Printf("Active Torrents: %d\n", stats.ActiveTorrents)
fmt.Printf("Used Space: %d bytes\n", stats.UsedSpace)
```

### Adding an RSS Feed

```go
req := models.AddRSSRequest{
    URL:  "https://example.com/feed.rss",
    Name: "My RSS Feed",
}

feed, err := client.General.AddRSS(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Added RSS feed: %s (ID: %d)\n", feed.Name, feed.ID)
```

### Searching Torrents

```go
torrents, err := client.General.SearchTorrents("ubuntu")
if err != nil {
    log.Fatal(err)
}

for _, t := range torrents {
    fmt.Printf("Found: %s (Size: %d)\n", t.Name, t.Size)
}
```

### Using Search API

```go
// Search by IMDB ID
torrents, err := client.Search.GetTorrent("imdb", "tt1234567")
if err != nil {
    log.Fatal(err)
}

// Get metadata
meta, err := client.Search.GetMeta("imdb", "tt1234567")
if err != nil {
    log.Fatal(err)
}
```

### Controlling Downloads

```go
import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

// Pause a torrent
err := client.General.ControlActiveTorrent(torrentID, constants.ControlActiveOperationPause)

// Resume a torrent
err = client.General.ControlActiveTorrent(torrentID, constants.ControlActiveOperationResume)

// Delete a torrent
err = client.General.ControlActiveTorrent(torrentID, constants.ControlActiveOperationDelete)

// Control usenet downloads
err = client.General.ControlUsenetDownload(usenetID, constants.ControlUsenetOperationPause)

// Control web downloads
err = client.General.ControlWebDownload(webID, constants.ControlWebDownloadOperationPause)
```

### Cloud Integration

```go
// Authorize Google Drive (code from OAuth flow)
err := client.General.AuthorizeGoogleDrive(authCode)

// Authorize Dropbox
err = client.General.AuthorizeDropbox(authCode)

// Check integration jobs
jobs, err := client.General.GetIntegrationJobs()
for _, job := range jobs {
    fmt.Printf("Job: %s - %s (%.2f%%)\n", job.FileName, job.Status, job.Progress*100)
}
```

## Environment Variables

The SDK respects the following environment variables:

- `TORBOX_API_KEY` - Your TorBox API key (required)
- `LOG_LEVEL` - Logging level (debug, info, warn, error)

## Architecture

The SDK follows a service-oriented architecture:

- **General Service** (`client.General`) - Core torrent, usenet, web downloads, user, RSS, and integration operations
- **Search Service** (`client.Search`) - Torrent search and metadata lookup

All services share a common HTTP client with:
- Automatic authentication via API key
- Retry logic with exponential backoff
- Rate limiting detection and handling
- Structured logging with zerolog

## Error Handling

All methods return errors that include context from the API:

```go
torrent, err := client.General.GetTorrentInfo(hash)
if err != nil {
    // Error messages include API error details
    log.Printf("Failed to get torrent info: %v", err)
}
```

## API Coverage

Based on the TorBox API v1 specification:

### Torrents
- ✅ `/api/torrents/createtorrent` - CreateTorrent
- ✅ `/api/torrents/controltorrent` - ControlActiveTorrent
- ✅ `/api/torrents/controlqueued` - ControlQueuedTorrent
- ✅ `/api/torrents/requestdl` - GetDownloadUrl
- ✅ `/api/torrents/mylist` - GetActiveTorrents
- ✅ `/api/torrents/checkcached` - CheckCached
- ✅ `/api/torrents/storesearch` - StoreSearch
- ✅ `/api/torrents/search` - SearchTorrents
- ✅ `/api/torrents/exportdata` - ExportData
- ✅ `/api/torrents/torrentinfo` - GetTorrentInfo
- ✅ `/api/torrents/getqueued` - GetQueuedTorrents

### Usenet
- ✅ `/api/usenet/createusenetdownload` - CreateUsenetDownload
- ✅ `/api/usenet/controlusenetdownload` - ControlUsenetDownload
- ✅ `/api/usenet/requestdl` - GetUsenetDownloadUrl
- ✅ `/api/usenet/mylist` - GetUsenetList
- ✅ `/api/usenet/checkcached` - CheckUsenetCached

### Web Downloads
- ✅ `/api/webdl/createwebdownload` - CreateWebDownload
- ✅ `/api/webdl/controlwebdownload` - ControlWebDownload

### User
- ✅ `/api/user/refreshtoken` - RefreshToken
- ✅ `/api/user/me` - GetUser
- ✅ `/api/user/addreferral` - AddReferral

### Notifications
- ✅ `/api/notifications/rss` - GetRSSNotifications
- ✅ `/api/notifications/mynotifications` - GetNotifications
- ✅ `/api/notifications/clear` - ClearNotifications

### RSS
- ✅ `/api/rss/addrss` - AddRSS
- ✅ `/api/rss/controlrss` - ControlRSS
- ✅ `/api/rss/modifyrss` - ModifyRSS

### Integration
- ✅ `/api/integration/googledrive` - AuthorizeGoogleDrive
- ✅ `/api/integration/dropbox` - AuthorizeDropbox
- ✅ `/api/integration/onedrive` - AuthorizeOneDrive
- ✅ `/api/integration/gofile` - AuthorizeGofile
- ✅ `/api/integration/1fichier` - Authorize1Fichier
- ✅ `/api/integration/jobs` - GetIntegrationJobs

### Stats
- ✅ `/api/stats` - GetStats

### Search API
- ✅ `/torrents/{type}:{id}` - GetTorrent
- ✅ `/meta/{type}:{id}` - GetMeta

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Acknowledgments

- TorBox API documentation
- Built with [zerolog](https://github.com/rs/zerolog) for structured logging
- JSON unmarshaling with [marshmallow](https://github.com/perimeterx/marshmallow)
