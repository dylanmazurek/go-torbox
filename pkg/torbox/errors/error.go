package errors

import "errors"

var (
	ErrServerError           = errors.New("server error")
	ErrDownloadAlreadyQueued = errors.New("download already queued")
	ErrInvalidMagnetLink     = errors.New("invalid magnet link")
)
