package monitor

import (
	"encoding/json"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"time"
)

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
	probe          probe.Probe
	previousProbes []*probe.ProbeResult
}

func NewMonitor(name string, interval time.Duration, p probe.Probe) *Monitor {
	return &Monitor{
		Name:           name,
		Interval:       interval,
		probe:          p,
		previousProbes: make([]*probe.ProbeResult, 0),
	}
}

func NewMonitorFromRecord(record monitorRecord) (*Monitor, error) {

	httpProbe := probe.HttpProbe{}

	err := json.Unmarshal([]byte(record.Definition), &httpProbe)

	if err != nil {
		return nil, err
	}

	m := &Monitor{}
	m.Name = record.Name
	m.Interval = record.Interval
	m.probe = &httpProbe
	return m, nil
}

func (m *Monitor) Status() Status {

	if len(m.previousProbes) == 0 {
		return StatusUnknown
	}

	if m.previousProbes[len(m.previousProbes)-1].Succeeded == probe.ExecutionSucceeded {
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

func (m *Monitor) GetPreviousProbes() []*probe.ProbeResult {
	return m.previousProbes
}

func (m *Monitor) IsSameAs(other *Monitor) bool {
	return m.Name == other.Name && m.Interval == other.Interval && m.probe.Target() == other.probe.Target()
}
