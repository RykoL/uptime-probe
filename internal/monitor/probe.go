package monitor

type ProbeStatus = bool

const (
	StatusSucceed = true
	StatusFailed  = false
)

type ProbeResult struct {
	Status ProbeStatus
}

type Probe interface {
	Execute() (*ProbeResult, error)
}
