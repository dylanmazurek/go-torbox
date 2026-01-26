package constants

import "errors"

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
	var opStr = string(data)
	o.FromString(opStr)

	return errors.New("unable to convert to control queued torrent operation")
}
