package model

import "time"

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
