package model

import (
	"slices"
	"time"
)

type Monitor struct {
	Id      int
	Name    string
	Results []ProbeResult
}

func (m *Monitor) Status() string {
	if len(m.Results) == 0 {
		return "Unknown"
	}
	if m.Results[len(m.Results)-1].Success {
		return "Up"
	}

	return "Down"
}

func (m *Monitor) OldestProbeResult() ProbeResult {
	return slices.MinFunc(m.Results, func(a, b ProbeResult) int {
		return a.Timestamp.Compare(b.Timestamp)
	})
}

func (m *Monitor) LatestProbeResult() ProbeResult {
	return slices.MaxFunc(m.Results, func(a, b ProbeResult) int {
		return a.Timestamp.Compare(b.Timestamp)
	})
}

type ProbeResult struct {
	Timestamp time.Time
	Success   bool
}

func (p *ProbeResult) Status() string {
	if p.Success {
		return "Up"
	}

	return "Down"
}
