# Copilot Instructions for go-torbox

## Project Overview

This is a Go SDK client library for the TorBox API, providing both general torrent management and search capabilities. The client follows a service-oriented architecture with two main API surfaces: General API for torrent CRUD operations and Search API for torrent discovery.

## Architecture Patterns

### Service Layer Structure
- **Main Client**: `pkg/torbox/client.go` - Factory that creates service instances with shared HTTP client
- **General Service**: `pkg/torbox/general/` - Handles torrent lifecycle (create, control, download URLs)
- **Search Service**: `pkg/torbox/search/` - Handles torrent discovery and metadata lookup
- **Models**: `pkg/torbox/models/` - Shared data structures across services

### HTTP Client Architecture
The project uses a custom transport layer pattern:
```go
// Authentication is handled via custom RoundTripper in auth.go
httpAuthClient := &http.Client{
    Transport: &addAuthHeaderTransport{
        T: &http.Transport{...},
        APIKey: clientOptions.apiKey,
    },
}
```

### Error Handling & Retry Logic
Both services implement sophisticated retry logic with exponential backoff:
- Rate limiting detection (429 responses) with `Retry-After` header respect
- Network error classification in `isRetryableNetworkError()` 
- 5xx server error retries with structured logging via zerolog
- Marshmallow JSON unmarshaling with unknown field warnings

## Key Conventions

### API Response Pattern
All API responses follow the BaseResponse structure:
```go
type BaseResponse struct {
    Success bool   `json:"success"`
    Error   string `json:"error"`
    Detail  string `json:"detail"`
    Data    any    `json:"data"`
}
```

### Request Building Pattern
Services use a common `newRequest()` method supporting multiple body types:
- `"json"` - Standard JSON marshaling
- `"form"` - URL-encoded form data
- `"file"` - Multipart file upload with custom Content-Type handling

### State Management
Torrent states are strongly typed via constants in `pkg/torbox/constants/state.go`:
- Processing states: `checkingResumeData`, `checking`, `metaDL`, `paused`
- Download states: `downloading`, `stalled (no seeds)`, `stalledDL`
- Upload states: `uploading`, `uploading (no peers)`
- Completion states: `completed`, `cached`

### Package Organization Rules
- **`internal/`**: Private utilities (crypto, form parsing, logging)
- **`pkg/torbox/`**: Public API client with sub-services
- **`pkg/magnet/`**: Standalone magnet link parser (no TorBox dependencies)
- **`pkg/torrent/`**: Bencode torrent file parser with crypto integration
- **`cmd/`**: CLI example/demo application

## Development Workflows

### Environment Setup
Requires `TORBOX_API_KEY` environment variable. Development setup uses:
- `.env` file for local development (gitignored)
- VS Code launch.json with `envFile` reference
- Logger respects `LOG_LEVEL` environment variable

### Testing Patterns
- Unit tests in `pkg/magnet/magnet_test.go` demonstrate table-driven test structure
- Test invalid inputs, edge cases, and expected outputs
- Use descriptive test names and structured assertions

### Context Propagation
Client initialization requires context, used primarily for logger setup:
```go
log := log.Ctx(ctx).With().Str("component", "torbox").Logger()
```

## Integration Points

### External Dependencies
- **marshmallow**: Flexible JSON unmarshaling with unknown field detection
- **zerolog**: Structured logging with custom console writer
- **zeebo/bencode**: Torrent file parsing
- **go-pretty**: Table rendering for CLI output

### API Endpoints
- **General API**: `https://api.torbox.app/v1` - Main torrent operations
- **Search API**: `https://search-api.torbox.app` - Search and metadata

### Torrent Lifecycle
1. **Creation**: Via magnet link or .torrent file upload (`CreateTorrent`)
2. **State Management**: Active vs Queued with different control operations
3. **Download**: Get signed download URLs for specific files
4. **Control**: Unified `ControlAnyTorrent` handles routing to active/queued APIs

## Key Files for Context
- `pkg/torbox/client.go` - Entry point and service initialization
- `pkg/torbox/general/service.go` - Core HTTP handling and retry logic
- `pkg/torbox/constants/` - API paths, states, and operation enums
- `pkg/torbox/models/torrent.go` - Complex JSON unmarshaling with embedded structs
- `cmd/main.go` - Complete client usage example