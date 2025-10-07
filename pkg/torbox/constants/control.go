package constants

type ControlActiveOperation string

const (
	ControlActiveOperationReannounce ControlActiveOperation = "reannounce"
	ControlActiveOperationDelete     ControlActiveOperation = "delete"
	ControlActiveOperationResume     ControlActiveOperation = "resume"
	ControlActiveOperationPause      ControlActiveOperation = "pause"
)

type ControlQueuedOperation string

const (
	ControlQueuedOperationDelete ControlQueuedOperation = "delete"
	ControlQueuedOperationStart  ControlQueuedOperation = "start"
)
