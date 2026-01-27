package constants

import (
	"errors"
	"strconv"
	"strings"
)

type ControlQueuedTorrentOperation int

const (
	ControlQueuedOperationDelete ControlQueuedTorrentOperation = iota
	ControlQueuedOperationStart
)

var controlQueuedOperationStrings = []string{
	"delete",
	"start",
}

func (o ControlQueuedTorrentOperation) String() string {
	return controlQueuedOperationStrings[o]
}

func (o *ControlQueuedTorrentOperation) FromString(opStr string) error {
	for i, str := range controlQueuedOperationStrings {
		if opStr == str {
			*o = ControlQueuedTorrentOperation(i)
			return nil
		}
	}

	return errors.New("unable to convert to control queued torrent operation")
}

func (o *ControlQueuedTorrentOperation) UnmarshalJSON(data []byte) error {
	dataStr := strings.Trim(string(data), "\"")

	if idx, err := strconv.Atoi(dataStr); err == nil {
		if idx >= 0 && idx < len(controlQueuedOperationStrings) {
			*o = ControlQueuedTorrentOperation(idx)
			return nil
		}
	}

	if err := o.FromString(dataStr); err == nil {
		return nil
	}

	return errors.New("unable to convert to control queued torrent operation")
}
