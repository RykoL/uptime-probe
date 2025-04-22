package monitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NoOpProbe struct {
}

func (p *NoOpProbe) Execute() (*ProbeResult, error) {
	return &ProbeResult{Status: true}, nil
}

func TestCreatesMonitorWithName(t *testing.T) {
	monitor := NewMonitor("Some Monitor", &NoOpProbe{})
	assert.Equal(t, monitor.Name, "Some Monitor")
}

func TestReturnsThatMonitorIsUpWhenLatestResultIsSuccessful(t *testing.T) {
	monitor := Monitor{Name: "", historicalData: []ProbeResult{
		{Status: StatusSucceed},
	}}

	assert.Equal(t, Status(StatusUp), monitor.Status())
}

func TestReturnsDownWhenLastResultIsAFailure(t *testing.T) {
	monitor := Monitor{Name: "", historicalData: []ProbeResult{
		{Status: StatusFailed},
	}}

	assert.Equal(t, Status(StatusDown), monitor.Status())
}

func TestReturnsUnknownWhenNoProbeHasExecuted(t *testing.T) {
	monitor := NewMonitor("Some", &NoOpProbe{})

	assert.Equal(t, Status(StatusUnknown), monitor.Status())
}
