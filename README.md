# go-torbox

A comprehensive Go SDK client library for the TorBox API, providing both general torrent management and search capabilities.

[![Go Reference](https://pkg.go.dev/badge/github.com/dylanmazurek/go-torbox.svg)](https://pkg.go.dev/github.com/dylanmazurek/go-torbox)
[![Go Version](https://img.shields.io/github/go-mod/go-version/dylanmazurek/go-torbox)](https://golang.org/dl/)

## Features

- 🚀 **Complete API Coverage**: Full implementation of TorBox General and Search APIs
- 🔄 **Automatic Retries**: Built-in retry logic with exponential backoff for network errors
- ⚡ **Rate Limiting**: Intelligent handling of API rate limits with `Retry-After` header support
- 📦 **Torrent Management**: Create, control, and download torrents with ease
- 🔍 **Search Capabilities**: Search torrents and fetch metadata
- 🧲 **Magnet Link Parsing**: Standalone magnet link parser
- 📄 **Torrent File Parsing**: Bencode torrent file parser with crypto support
- 📝 **Structured Logging**: Integration with zerolog for comprehensive logging
- 🔐 **Secure Authentication**: Custom transport layer with API key authentication

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
    "os"

    "github.com/dylanmazurek/go-torbox/pkg/torbox"
)

func main() {
    ctx := context.Background()

    // Get API key from environment
    apiKey := os.Getenv("TORBOX_API_KEY")
    if apiKey == "" {
        log.Fatal("TORBOX_API_KEY environment variable is not set")
    }

    // Create client
    client, err := torbox.New(ctx, torbox.WithAPIKey(apiKey))
    if err != nil {
        log.Fatal(err)
    }

    // Get active torrents
    torrents, err := client.General.GetActiveTorrents()
    if err != nil {
        log.Fatal(err)
    }

    for _, torrent := range torrents {
        fmt.Printf("ID: %d, Name: %s, Status: %s\n", 
            torrent.ID, torrent.Name, torrent.DownloadState)
    }
}
```

## Architecture

The client follows a service-oriented architecture with two main API surfaces:

### Main Client

The main client acts as a factory that creates service instances with a shared HTTP client:

```go
type Client struct {
    General *general.GeneralService
    Search  *search.SearchService
}
```

### General Service

Handles the torrent lifecycle:
- Create torrents via magnet link or .torrent file
- Control torrent operations (pause, resume, delete)
- Get download URLs for files
- Manage active and queued torrents

### Search Service

Handles torrent discovery:
- Search for torrents
- Fetch torrent metadata

## Usage Examples

### Creating a Torrent from Magnet Link

```go
import (
    "github.com/dylanmazurek/go-torbox/pkg/magnet"
    "github.com/dylanmazurek/go-torbox/pkg/torbox/models"
    "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
)

magnetLink, err := magnet.NewMagnet("magnet:?xt=urn:btih:...")
if err != nil {
    log.Fatal(err)
}

// Create with optional settings
seedSetting := constants.Auto
allowZip := true
name := "My Torrent"
asQueued := false

request := models.CreateTorrentRequest{
    Magnet:   magnetLink,
    Seed:     &seedSetting,
    AllowZip: &allowZip,
    Name:     &name,
    AsQueued: &asQueued,
}

torrent, err := client.General.CreateTorrent(request)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Created torrent: %s\n", torrent.Name)
```

### Creating a Torrent from File

```go
import (
    "os"
    "github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

fileData, err := os.ReadFile("path/to/file.torrent")
if err != nil {
    log.Fatal(err)
}

request := models.CreateTorrentRequest{
    File: fileData,
}

torrent, err := client.General.CreateTorrent(request)
if err != nil {
    log.Fatal(err)
}
```

### Controlling Torrents

```go
import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

// Pause an active torrent
err := client.General.ControlActiveTorrent(torrentID, constants.ControlActiveOperationPause)

// Resume an active torrent
err = client.General.ControlActiveTorrent(torrentID, constants.ControlActiveOperationResume)

// Control any torrent (automatically routes to active or queued API)
err = client.General.ControlAnyTorrent(torrentID, "pause")
```

### Getting Download URLs

```go
downloadURL, err := client.General.GetDownloadUrl(torrentID, fileID)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Download URL: %s\n", *downloadURL)
```

### Getting Queued Torrents

```go
queuedTorrents, err := client.General.GetQueuedTorrents()
if err != nil {
    log.Fatal(err)
}

for _, queued := range queuedTorrents {
    fmt.Printf("Queued ID: %d, Name: %s\n", queued.ID, queued.Name)
}
```

### Searching for Torrents

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

### Parsing Torrent Files

```go
import "github.com/dylanmazurek/go-torbox/pkg/torrent"

// Parse from file
torrentInfo, err := torrent.ParseFromFile("path/to/file.torrent")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("InfoHash: %s\n", torrentInfo.InfoHash)
fmt.Printf("Files: %d\n", len(torrentInfo.Files))

// Parse from reader
file, err := os.Open("path/to/file.torrent")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

torrentInfo, err = torrent.Parse(file)
```

### Parsing Magnet Links

```go
import "github.com/dylanmazurek/go-torbox/pkg/magnet"

magnetLink := "magnet:?xt=urn:btih:1234567890abcdef&dn=example"
magnet, err := magnet.NewMagnet(magnetLink)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Hash: %s\n", magnet.Hash)
fmt.Printf("Display Name: %s\n", magnet.DisplayName)
```

## API Reference

### Client Options

```go
// Create client with API key
client, err := torbox.New(ctx, torbox.WithAPIKey("your-api-key"))
```

### General Service Methods

| Method | Description |
|--------|-------------|
| `GetActiveTorrents()` | Retrieve all active torrents |
| `GetQueuedTorrents()` | Retrieve all queued torrents |
| `CreateTorrent(request)` | Create a new torrent from magnet link or file |
| `GetDownloadUrl(torrentId, fileId)` | Get download URL for a specific file |
| `ControlActiveTorrent(id, operation)` | Control an active torrent (pause, resume, etc.) |
| `ControlQueuedTorrent(id, operation)` | Control a queued torrent |
| `ControlAnyTorrent(id, operation)` | Control any torrent (auto-routes to active/queued) |

### Search Service Methods

| Method | Description |
|--------|-------------|
| `GetTorrent(idType, id)` | Search for torrents by ID type (imdb, tmdb, etc.) |
| `GetMeta(idType, id)` | Get torrent metadata by ID type |

### Torrent States

The library defines the following torrent states (in `pkg/torbox/constants/state.go`):

- Processing: `checkingResumeData`, `checking`, `metaDL`, `paused`
- Downloading: `downloading`, `stalled (no seeds)`, `stalledDL`
- Uploading: `uploading`, `uploading (no peers)`
- Completion: `completed`, `cached`

### Seed Settings

Available seed settings when creating torrents:

- `constants.Auto` (1) - Automatically decide seeding behavior
- `constants.Seed` (2) - Force seeding
- `constants.NoSeed` (3) - Disable seeding

### Control Operations

Available operations for torrent control:

**Active Torrents:**
- `ControlActiveOperationPause` - Pause the torrent
- `ControlActiveOperationResume` - Resume the torrent
- `ControlActiveOperationReannounce` - Reannounce to trackers
- `ControlActiveOperationDelete` - Delete the torrent

**Queued Torrents:**
- `ControlQueuedOperationStart` - Start the queued torrent
- `ControlQueuedOperationDelete` - Delete the queued torrent

## Advanced Features

### Retry Logic

The client automatically retries failed requests with exponential backoff:

- Network errors: Retries up to 3 times with exponential delay
- Rate limiting (429): Respects `Retry-After` header or uses exponential backoff
- Server errors (5xx): Automatic retry with backoff
- Configurable timeout (60 seconds default)

### Logging

The client uses structured logging with zerolog:

```go
import "github.com/dylanmazurek/go-torbox/internal/logger"

ctx := context.Background()
log.Logger = logger.New(ctx)
```

Set log level via environment variable:
```bash
export LOG_LEVEL=debug
```

### HTTP Client Configuration

The client uses a custom transport layer with:
- Connection pooling (10 max idle connections)
- Keep-alive support
- 60-second timeout for file operations
- Automatic authentication header injection

## Package Structure

```
pkg/
├── torbox/              # Main client package
│   ├── client.go        # Client factory
│   ├── general/         # General API service
│   ├── search/          # Search API service
│   ├── models/          # Request/response models
│   └── constants/       # API constants and enums
├── magnet/              # Magnet link parser
└── torrent/             # Torrent file parser

internal/
├── crypto/              # Cryptographic utilities
├── logger/              # Logging setup
└── form/                # Form encoding utilities

cmd/
└── main.go              # Example CLI application
```

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `TORBOX_API_KEY` | Your TorBox API key | Yes |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | No |

## Development

### Prerequisites

- Go 1.24.2 or later

### Running Tests

```bash
go test ./...
```

### Running the Example CLI

```bash
export TORBOX_API_KEY=your-api-key
go run cmd/main.go
```

### Building

```bash
go build ./...
```

## API Endpoints

The client interacts with the following TorBox API endpoints:

- **General API**: `https://api.torbox.app/v1`
- **Search API**: `https://search-api.torbox.app`

## Error Handling

The library uses Go's standard error handling. All methods return errors that should be checked:

```go
torrents, err := client.General.GetActiveTorrents()
if err != nil {
    // Handle error
    log.Printf("Failed to get torrents: %v", err)
    return
}
```

API errors include details from the TorBox API response:

```go
type BaseResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
    Detail  string `json:"detail"`
    Data    any    `json:"data"`
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is provided as-is. Please refer to the repository for license information.

## Links

- [TorBox Website](https://torbox.app)
- [TorBox API Documentation](https://torbox.app/api-docs)
- [Go Package Documentation](https://pkg.go.dev/github.com/dylanmazurek/go-torbox)

## Acknowledgments

This library uses the following open-source packages:

- [zerolog](https://github.com/rs/zerolog) - Fast and simple logging
- [marshmallow](https://github.com/perimeterx/marshmallow) - Flexible JSON unmarshaling
- [bencode](https://github.com/zeebo/bencode) - Bencode encoder/decoder
- [go-pretty](https://github.com/jedib0t/go-pretty) - Pretty table rendering
