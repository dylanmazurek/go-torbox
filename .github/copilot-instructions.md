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

## Code Style & Formatting

### Go Standards
- Follow standard Go formatting with `gofmt`
- Use `go vet` for static analysis
- Organize imports in groups: standard library, external packages, internal packages
- Keep functions focused and small (prefer composition over large functions)

### Naming Conventions
- **Types**: PascalCase for exported types, camelCase for unexported
- **Constants**: PascalCase with descriptive prefixes (e.g., `ControlActiveOperation`, `PATH_TORRENTS_GET_ACTIVE`)
- **Packages**: Lowercase, single word when possible (e.g., `general`, `search`, `models`)
- **Interfaces**: Use "-er" suffix for single-method interfaces
- **Errors**: Prefix with `Err` for package-level errors (e.g., `ErrServerError`)

### File Organization
- Group related functionality in separate files (e.g., `active.go`, `queued.go`, `torrents.go`)
- Place constants in dedicated files under `pkg/torbox/constants/`
- Keep models in `pkg/torbox/models/` with one type per file when complex
- Service methods organized by entity type (active torrents, queued torrents, etc.)

## Build & Development Commands

### Running the Project
```bash
# Set up environment
export TORBOX_API_KEY="your-api-key"
# Or use .env file (gitignored)

# Run the example CLI
go run cmd/main.go

# Build the CLI
go build -o torbox-cli cmd/main.go
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for specific package
go test ./pkg/magnet/...

# Run with race detection
go test -race ./...
```

### Code Quality
```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Tidy dependencies
go mod tidy
```

## Common Development Tasks

### Adding a New Service Method
1. Define request/response models in `pkg/torbox/models/`
2. Add API path constant to `pkg/torbox/constants/constants.go`
3. Implement method in appropriate service (`general` or `search`)
4. Use `newRequest()` helper with appropriate body type
5. Handle response with `do()` or `doWithRetry()`
6. Add example usage to `cmd/main.go` if demonstrating new functionality

### Adding New Constants
- **API paths**: Add to `pkg/torbox/constants/constants.go`
- **States**: Add to `pkg/torbox/constants/state.go`
- **Operations**: Add to `pkg/torbox/constants/control.go`
- Use typed constants (e.g., `type ControlActiveOperation string`) for type safety

### Error Handling Best Practices
- Define package-level errors in `pkg/torbox/errors/error.go`
- Use `fmt.Errorf` with context for runtime errors
- Check `BaseResponse.Success` before accessing data
- Provide detailed error messages combining `Error` and `Detail` fields
- Let retry logic handle transient failures automatically

## Key Files for Context
- `pkg/torbox/client.go` - Entry point and service initialization
- `pkg/torbox/general/service.go` - Core HTTP handling and retry logic
- `pkg/torbox/constants/` - API paths, states, and operation enums
- `pkg/torbox/models/torrent.go` - Complex JSON unmarshaling with embedded structs
- `cmd/main.go` - Complete client usage example
- `pkg/torbox/errors/error.go` - Package-level error definitions