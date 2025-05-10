package monitor

import (
	"context"
	"encoding/json"
	"fmt"
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
	Id             int
	Name           string
	Interval       time.Duration
	probe          probe.Probe
	previousProbes []*probe.ProbeResult
	isRunning      bool
}

func NewMonitor(name string, interval time.Duration, p probe.Probe) *Monitor {
	return &Monitor{
		Id:             -1,
		Name:           name,
		Interval:       interval,
		probe:          p,
		previousProbes: make([]*probe.ProbeResult, 0),
	}
}

func (m *Monitor) Start(ctx context.Context, repository Repository) {
	ticker := time.NewTicker(m.Interval)
	defer ticker.Stop()

	fmt.Printf("Starting monitor %s\n", m.Name)

	m.isRunning = true

loopExit:
	for {
		err, result := m.Probe()

		err = repository.RecordProbeResult(ctx, m.Id, result)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			break loopExit
		}
	}

	m.isRunning = false
	fmt.Printf("Monitor %s stopped \n", m.Name)
}

func NewMonitorFromRecord(record monitorRecord) (*Monitor, error) {

	httpProbe := probe.HttpProbe{}

	err := json.Unmarshal([]byte(record.Definition), &httpProbe)

	if err != nil {
		return nil, err
	}

	m := &Monitor{
		Id:       record.Id,
		Name:     record.Name,
		Interval: record.Interval,
		probe:    &httpProbe,
	}
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

func (m *Monitor) Probe() (error, *probe.ProbeResult) {

	fmt.Printf("Executing probe for %s\n", m.Name)
	result, err := m.probe.Execute()

	if err != nil {
		return err, nil
	}

	result.TimeStamp = time.Now()

	return nil, result
}

func (m *Monitor) GetPreviousProbes() []*probe.ProbeResult {
	return m.previousProbes
}

func (m *Monitor) Equals(other *Monitor) bool {
	return m.Name == other.Name && m.Interval == other.Interval && m.probe.Target() == other.probe.Target()
}
