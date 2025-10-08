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

type ControlUsenetOperation string

const (
	ControlUsenetOperationDelete ControlUsenetOperation = "delete"
	ControlUsenetOperationResume ControlUsenetOperation = "resume"
	ControlUsenetOperationPause  ControlUsenetOperation = "pause"
)

type ControlWebDownloadOperation string

const (
	ControlWebDownloadOperationDelete ControlWebDownloadOperation = "delete"
	ControlWebDownloadOperationResume ControlWebDownloadOperation = "resume"
	ControlWebDownloadOperationPause  ControlWebDownloadOperation = "pause"
)

type ControlRSSOperation string

const (
	ControlRSSOperationDelete ControlRSSOperation = "delete"
	ControlRSSOperationPause  ControlRSSOperation = "pause"
	ControlRSSOperationResume ControlRSSOperation = "resume"
)
