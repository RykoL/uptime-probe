package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"time"
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

	m.isRunning = true

loopExit:
	for {
		err, result := m.Probe()

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

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

func (m *Monitor) Probe() (error, *probe.ProbeResult) {

	result, err := m.probe.Execute()

	if err != nil {
		return err, nil
	}

	result.TimeStamp = time.Now()

	return nil, result
}

func (m *Monitor) Equals(other *Monitor) bool {
	return m.Name == other.Name && m.Interval == other.Interval && m.probe.Target() == other.probe.Target()
}
