package monitor

import "time"

type ExecutionStatus = bool

const (
	ExecutionSucceeded = true
	ExecutionFailed    = false
)

type ProbeResult struct {
	TimeStamp time.Time
	Succeeded ExecutionStatus
}

type Probe interface {
	Execute() (*ProbeResult, error)
}
