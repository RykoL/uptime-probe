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
	previousProbes []*ProbeResult
}

func NewMonitor(name string, interval time.Duration, probe Probe) *Monitor {
	return &Monitor{
		Name:           name,
		Interval:       interval,
		probe:          probe,
		previousProbes: make([]*ProbeResult, 0),
	}
}

func (m *Monitor) Status() Status {

	if len(m.previousProbes) == 0 {
		return StatusUnknown
	}

	if m.previousProbes[len(m.previousProbes)-1].Succeeded == ExecutionSucceeded {
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
	if len(m.previousProbes) == 0 {
		return true
	}

	lastExecution := m.previousProbes[len(m.previousProbes)-1].TimeStamp

	if time.Now().Add(-m.Interval).After(lastExecution) {
		return true
	}

	return false
}

func (m *Monitor) Probe() error {

	result, err := m.probe.Execute()

	if err != nil {
		return err
	}

	result.TimeStamp = time.Now()

	m.previousProbes = append(m.previousProbes, result)

	return nil
}

func (m *Monitor) GetPreviousProbes() []*ProbeResult {
	return m.previousProbes
}
