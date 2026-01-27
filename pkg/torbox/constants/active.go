package constants

import (
	"errors"
	"strconv"
	"strings"
)

type ControlActiveTorrentOperation int

const (
	ControlActiveOperationReannounce ControlActiveTorrentOperation = iota
	ControlActiveOperationDelete
	ControlActiveOperationResume
	ControlActiveOperationPause
)

var controlActiveOperationStrings = []string{
	"reannounce",
	"delete",
	"resume",
	"pause",
}

func (o ControlActiveTorrentOperation) String() string {
	return controlActiveOperationStrings[o]
}

func (o *ControlActiveTorrentOperation) FromString(opStr string) error {
	for i, str := range controlActiveOperationStrings {
		if opStr == str {
			*o = ControlActiveTorrentOperation(i)
			return nil
		}
	}

	return errors.New("unable to convert to control active torrent operation")
}

func (o *ControlActiveTorrentOperation) UnmarshalJSON(data []byte) error {
	dataStr := strings.Trim(string(data), "\"")

	if idx, err := strconv.Atoi(dataStr); err == nil {
		if idx >= 0 && idx < len(controlActiveOperationStrings) {
			*o = ControlActiveTorrentOperation(idx)
			return nil
		}
	}

	if err := o.FromString(dataStr); err == nil {
		return nil
	}

	return errors.New("unable to convert to control active torrent operation")
}
