package monitor

import "time"

type ProbeStatus = bool

const (
	StatusSucceed = true
	StatusFailed  = false
)

type ProbeResult struct {
	TimeStamp time.Time
	Status    ProbeStatus
}

type Probe interface {
	Execute() (*ProbeResult, error)
}
