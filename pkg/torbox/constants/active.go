package constants

import (
	"errors"
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
	var opStr = string(data)
	o.FromString(opStr)

	return errors.New("unable to convert to control active torrent operation")
}
