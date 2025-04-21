package monitor

import "github.com/RykoL/uptime-probe/internal/probe"

type Status int

const (
	StatusDown = iota
	StatusUp
	StatusPending
	StatusUnknown
)

type Monitor struct {
	Name           string
	historicalData []probe.Result
}

func NewMonitor(name string) Monitor {
	return Monitor{
		Name:           name,
		historicalData: make([]probe.Result, 0),
	}
}

func (m *Monitor) Status() Status {

	if len(m.historicalData) == 0 {
		return StatusUnknown
	}

	if m.historicalData[len(m.historicalData)-1].Status == probe.StatusSucceed {
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
