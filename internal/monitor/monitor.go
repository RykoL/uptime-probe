package monitor

import "time"

type Status int

const (
	StatusDown = iota
	StatusUp
	StatusPending
	StatusUnknown
)

type Monitor struct {
	Name           string
	Interval       time.Duration
	probe          Probe
	historicalData []ProbeResult
}

func NewMonitor(name string, interval time.Duration, probe Probe) *Monitor {
	return &Monitor{
		Name:           name,
		Interval:       interval,
		probe:          probe,
		historicalData: make([]ProbeResult, 0),
	}
}

func (m *Monitor) Status() Status {

	if len(m.historicalData) == 0 {
		return StatusUnknown
	}

	if m.historicalData[len(m.historicalData)-1].Status == StatusSucceed {
		return StatusUp
	}

	return StatusDown
}

func (s Status) String() string {
	switch s {
	case StatusUp:
		return "Up"
	case StatusUnknown:
		return "Unknown"
	case StatusPending:
		return "Pending"
	case StatusDown:
		return "Down"
	default:
		return "Error: This status doesn't exists"
	}
}

func (m *Monitor) ShouldExecuteProbe() bool {
	if len(m.historicalData) == 0 {
		return true
	}

	lastExecution := m.historicalData[len(m.historicalData)-1].TimeStamp

	if time.Now().Add(-m.Interval).After(lastExecution) {
		return true
	}

	return false
}

func (m *Monitor) Probe() error {
	return nil
}
