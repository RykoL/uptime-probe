package monitor

type Status int

const (
	StatusDown = iota
	StatusUp
	StatusPending
	StatusUnknown
)

type Monitor struct {
	Name           string
	Probe          Probe
	historicalData []ProbeResult
}

func NewMonitor(name string, probe Probe) *Monitor {
	return &Monitor{
		Name:           name,
		Probe:          probe,
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
